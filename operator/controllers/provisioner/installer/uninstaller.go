// Copyright 2020 NetApp, Inc. All Rights Reserved.

package installer

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/policy/v1beta1"
	v12 "k8s.io/api/rbac/v1"
	v1beta12 "k8s.io/api/storage/v1beta1"

	"github.com/netapp/trident/cli/cmd"
	k8sclient "github.com/netapp/trident/cli/k8s_client"
)

func (i *Installer) isTridentInstalled() (installed bool, namespace string, err error) {
	return i.client.CheckDeploymentExistsByLabel(TridentLegacyLabel, true)
}

func (i *Installer) isPreviewCSITridentInstalled() (installed bool, namespace string, err error) {
	return i.client.CheckStatefulSetExistsByLabel(TridentCSILabel, true)
}

func (i *Installer) isCSITridentInstalled() (installed bool, namespace string, err error) {
	return i.client.CheckDeploymentExistsByLabel(TridentCSILabel, true)
}

func (i *Installer) discoverTrident() (legacy, csi, csiPreview bool, err error) {

	// Check if legacy Trident is installed
	if legacy, _, err = i.isTridentInstalled(); err != nil {
		err = fmt.Errorf("could not check if legacy Trident is installed; %v", err)
		return
	}

	// Check if preview CSI Trident is installed
	if csiPreview, _, err = i.isPreviewCSITridentInstalled(); err != nil {
		err = fmt.Errorf("could not check if preview CSI Trident is installed; %v", err)
		return
	}

	// Check if CSI Trident is installed
	if csi, _, err = i.isCSITridentInstalled(); err != nil {
		err = fmt.Errorf("could not check if CSI Trident is installed; %v", err)
		return
	}

	return
}

func (i *Installer) UninstallTrident() error {

	// 1. preview CSI Trident --> uninstall preview CSI Trident
	// 2. preview CSI Trident & legacy Trident --> uninstall preview CSI Trident
	// 3. CSI Trident --> uninstall CSI Trident
	// 4. legacy Trident --> uninstall legacy Trident
	//
	// if csiPreview, uninstall csiPreview
	// else if csi, uninstall csi
	// else if legacy, uninstall legacy

	legacyTridentInstalled, _, csiPreviewTridentInstalled, err := i.discoverTrident()
	if err != nil {
		return err
	}

	if legacyTridentInstalled && csiPreviewTridentInstalled {
		log.Warning("Both legacy and CSI Trident are installed.  CSI Trident will be uninstalled, and " +
			"the uninstaller will run again to remove legacy Trident before running the Trident installer.")
	}

	// Set the global csi variable, which controls things like RBAC and app labels
	// Should not use csiPreviewTridentInstalled || csiTridentInstalled as it give false when CSI trident
	// installtion is deleted
	csi := !legacyTridentInstalled

	// Set the app labels (CSI takes precedence)
	if csi {
		appLabel = TridentCSILabel
		appLabelKey = TridentCSILabelKey
		appLabelValue = TridentCSILabelValue
	} else {
		appLabel = TridentLegacyLabel
		appLabelKey = TridentLegacyLabelKey
		appLabelValue = TridentLegacyLabelValue
	}

	// First handle the deployment (legacy, CSI) / statefulset (preview CSI)
	if csiPreviewTridentInstalled {
		if err := i.deleteTridentStatefulSet(); err != nil {
			return err
		}

	} else {
		if err := i.deleteTridentDeployment(); err != nil {
			return err
		}
	}

	// Next handle all the other common CSI components (daemonset, service).  Some/all of these may
	// not be present if uninstalling legacy Trident or preview CSI Trident, in which case we log
	// warnings only.

	if err := i.deleteTridentDaemonSet(); err != nil {
		return err
	}

	if err := i.deleteTridentService(); err != nil {
		return err
	}

	if err := i.deleteTridentSecret(); err != nil {
		return err
	}

	if err := i.deleteTridentCSIDriverCR(); err != nil {
		return err
	}

	if err := i.removeRBACObjects(log.InfoLevel); err != nil {
		return err
	}

	if err := i.deletePodSecurityPolicy(); err != nil {
		return err
	}

	log.Info("The uninstaller did not delete Trident's namespace in case it is going to be reused.")

	return nil
}

func (i *Installer) UninstallCSIPreviewTrident() error {
	appLabel = TridentCSILabel
	appLabelKey = TridentCSILabelKey
	appLabelValue = TridentCSILabelValue

	return i.deleteTridentStatefulSet()
}

func (i *Installer) UninstallLegacyTrident() error {
	appLabel = TridentLegacyLabel
	appLabelKey = TridentLegacyLabelKey
	appLabelValue = TridentLegacyLabelValue

	if err := i.deleteTridentDeployment(); err != nil {
		return err
	}

	return i.removeRBACObjects(log.InfoLevel)
}

func (i *Installer) removeRBACObjects(logLevel log.Level) error {

	logFunc := func(fields log.Fields) func(args ...interface{}) {
		if logLevel == log.DebugLevel {
			return log.WithFields(fields).Debug
		} else {
			return log.WithFields(fields).Info
		}
	}

	// Delete cluster role binding
	if err := i.deleteTridentClusterRoleBinding(); err != nil {
		return err
	}

	// Delete cluster role
	if err := i.deleteTridentClusterRole(); err != nil {
		return err
	}

	// Delete service account
	if err := i.deleteTridentServiceAccount(); err != nil {
		return err
	}

	// If OpenShift, remove Trident from security context constraint(s)
	if i.client.Flavor() == k8sclient.FlavorOpenShift {
		if csi {
			if err := i.client.RemoveTridentUserFromOpenShiftSCC("trident-csi", "privileged"); err != nil {
				log.WithField("error", err).Warning("Could not modify security context constraint.")
				return fmt.Errorf("could not modify security context constraint for trident-csi; err: %v", err)
			} else {
				logFunc(log.Fields{
					"scc":  "privileged",
					"user": "trident-csi",
				})("Removed security context constraint user.")
			}
		} else {
			if err := i.client.RemoveTridentUserFromOpenShiftSCC("trident", "anyuid"); err != nil {
				log.WithField("error", err).Warning("Could not modify security context constraint.")
				return fmt.Errorf("could not modify security context constraint for trident; err: %v", err)
			} else {
				logFunc(log.Fields{
					"scc":  "anyuid",
					"user": "trident",
				})("Removed security context constraint user.")
			}
		}
	}

	return nil
}

func (i *Installer) deleteTridentStatefulSet() error {

	// Delete Trident statefulSet
	if statefulSets, err := i.client.GetStatefulSetsByLabel(appLabel, true); err != nil {
		log.Errorf("Unable to get list of stateful sets by label %v", appLabel)
		return fmt.Errorf("unable to get list of stateful sets by label")
	} else if len(statefulSets) == 0 {
		log.WithFields(log.Fields{
			"label": appLabel,
			"error": err,
		}).Warn("Trident Stateful set not found.")
	} else {
		if len(statefulSets) == 1 {
			log.WithFields(log.Fields{
				"service":   statefulSets[0].Name,
				"namespace": statefulSets[0].Namespace,
			}).Info("Trident Stateful set found by label.")
		} else {
			log.Warnf("Multiple Stateful sets found matching label: %s; removing all.", appLabel)
		}

		if err = i.RemoveMultipleStatefulSets(statefulSets); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) RemoveMultipleStatefulSets(unwantedStatefulSets []appsv1.StatefulSet) error {
	var err error
	var anyError bool
	var undeletedStatefulSets []string

	if len(unwantedStatefulSets) > 0 {
		for _, statefulSetToRemove := range unwantedStatefulSets {
			// Delete the stateful set
			if err = i.client.DeleteStatefulSet(statefulSetToRemove.Name, statefulSetToRemove.Namespace); err != nil {
				log.WithFields(log.Fields{
					"statefulset": statefulSetToRemove.Name,
					"namespace":   statefulSetToRemove.Namespace,
					"label":       appLabel,
					"error":       err,
				}).Errorf("Could not delete Trident Stateful set.")

				anyError = true
				undeletedStatefulSets = append(undeletedStatefulSets, fmt.Sprintf("%v/%v", statefulSetToRemove.Namespace,
					statefulSetToRemove.Name))
			} else {
				log.Infof("Deleted Trident Stateful set: %v/%v", statefulSetToRemove.Namespace, statefulSetToRemove.Name)
			}
		}
	}

	if anyError {
		return fmt.Errorf("unable to delete Trident Stateful set(s): %v", undeletedStatefulSets)
	}

	return nil
}

func (i *Installer) deleteTridentDeployment() error {

	// Delete Trident deployments
	if deployments, err := i.client.GetDeploymentsByLabel(appLabel, true); err != nil {

		log.Errorf("Unable to get list of deployments by label %v", appLabel)
		return fmt.Errorf("unable to get list of deployments")

	} else if len(deployments) == 0 {
		log.WithFields(log.Fields{
			"label": appLabel,
			"error": err,
		}).Warn("Trident deployment not found.")
	} else {

		if len(deployments) == 1 {
			log.WithFields(log.Fields{
				"service":   deployments[0].Name,
				"namespace": deployments[0].Namespace,
			}).Info("Trident deployment found by label.")
		} else {
			log.Warnf("Multiple deployments found matching label: %s; removing all.", appLabel)
		}

		if err = i.RemoveMultipleDeployments(deployments); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) RemoveMultipleDeployments(unwantedDeployments []appsv1.Deployment) error {
	var err error
	var anyError bool
	var undeletedDeployments []string

	if len(unwantedDeployments) > 0 {
		for _, deploymentToRemove := range unwantedDeployments {
			// Delete the deployment
			if err = i.client.DeleteDeployment(deploymentToRemove.Name, deploymentToRemove.Namespace); err != nil {
				log.WithFields(log.Fields{
					"deployment": deploymentToRemove.Name,
					"namespace":  deploymentToRemove.Namespace,
					"label":      appLabel,
					"error":      err,
				}).Errorf("Could not delete Trident deployment.")

				anyError = true
				undeletedDeployments = append(undeletedDeployments, fmt.Sprintf("%v/%v", deploymentToRemove.Namespace,
					deploymentToRemove.Name))
			} else {
				log.Infof("Deleted Trident deployment: %v/%v", deploymentToRemove.Namespace, deploymentToRemove.Name)
			}
		}
	}

	if anyError {
		return fmt.Errorf("unable to delete Trident deployment(s): %v", undeletedDeployments)
	}

	return nil
}

func (i *Installer) deleteTridentDaemonSet() error {

	// Delete Trident daemonsets
	if daemonsets, err := i.client.GetDaemonSetsByLabel(TridentNodeLabel, true); err != nil {

		log.Errorf("Unable to get list of daemonset by label %v", TridentNodeLabel)
		return fmt.Errorf("unable to get list of daemonset")

	} else if len(daemonsets) == 0 {
		log.WithFields(log.Fields{
			"label": TridentNodeLabel,
			"error": err,
		}).Warning("Trident daemonset not found.")

	} else {
		if len(daemonsets) == 1 {
			log.WithFields(log.Fields{
				"service":   daemonsets[0].Name,
				"namespace": daemonsets[0].Namespace,
			}).Info("Trident daemonsets found by label.")
		} else {
			log.Warnf("Multiple daemonsets found matching label: %s; removing all.", TridentNodeLabel)
		}

		if err = i.RemoveMultipleDaemonSets(daemonsets); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) RemoveMultipleDaemonSets(unwantedDaemonsets []appsv1.DaemonSet) error {
	var err error
	var anyError bool
	var undeletedDaemonsets []string

	if len(unwantedDaemonsets) > 0 {
		for _, daemonsetToRemove := range unwantedDaemonsets {
			// Delete the daemonset
			if err = i.client.DeleteDaemonSet(daemonsetToRemove.Name, daemonsetToRemove.Namespace); err != nil {
				log.WithFields(log.Fields{
					"deployment": daemonsetToRemove.Name,
					"namespace":  daemonsetToRemove.Namespace,
					"label":      TridentNodeLabel,
					"error":      err,
				}).Warning("Could not delete Trident daemonset.")

				anyError = true
				undeletedDaemonsets = append(undeletedDaemonsets, fmt.Sprintf("%v/%v", daemonsetToRemove.Namespace,
					daemonsetToRemove.Name))
			} else {
				log.Infof("Deleted Trident daemonset: %v/%v", daemonsetToRemove.Namespace, daemonsetToRemove.Name)
			}
		}
	}

	if anyError {
		return fmt.Errorf("unable to delete Trident daemonset(s): %v", undeletedDaemonsets)
	}

	return nil
}

func (i *Installer) deleteTridentService() error {

	// Delete Trident services
	if services, err := i.client.GetServicesByLabel(appLabel, true); err != nil {
		log.Errorf("Unable to get list of services by label %v", appLabel)
		return fmt.Errorf("unable to get list of services")
	} else if len(services) == 0 {
		log.WithFields(log.Fields{
			"label": appLabel,
			"error": err,
		}).Warning("Trident service not found.")
	} else {
		if len(services) == 1 {
			log.WithFields(log.Fields{
				"service":   services[0].Name,
				"namespace": services[0].Namespace,
			}).Info("Trident service found by label.")
		} else {
			log.Warnf("Multiple services found matching label: %s; removing all.", appLabel)
		}

		if err = i.RemoveMultipleServices(services); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) RemoveMultipleServices(unwantedServices []v1.Service) error {
	var err error
	var anyError bool
	var undeletedServices []string

	if len(unwantedServices) > 0 {
		for _, serviceToRemove := range unwantedServices {
			// Delete the service
			if err = i.client.DeleteService(serviceToRemove.Name, serviceToRemove.Namespace); err != nil {
				log.WithFields(log.Fields{
					"service":   serviceToRemove.Name,
					"namespace": serviceToRemove.Namespace,
					"label":     appLabel,
					"error":     err,
				}).Warning("Could not delete Trident service.")

				anyError = true
				undeletedServices = append(undeletedServices, fmt.Sprintf("%v/%v", serviceToRemove.Namespace,
					serviceToRemove.Name))
			} else {
				log.Infof("Deleted Trident service: %v/%v", serviceToRemove.Namespace, serviceToRemove.Name)
			}
		}
	}

	if anyError {
		return fmt.Errorf("unable to delete Trident service(s): %v", undeletedServices)
	}

	return nil
}

func (i *Installer) deleteTridentSecret() error {

	if secrets, err := i.client.GetSecretsByLabel(appLabel, false); err != nil {
		log.Errorf("Unable to get list of Secrets by label %v", appLabel)
		return fmt.Errorf("unable to get list of Secrets by label")
	} else if len(secrets) == 0 {
		log.WithFields(log.Fields{
			"label": appLabel,
			"error": err,
		}).Warning("Trident secret not found.")
	} else {
		if len(secrets) == 1 {
			log.WithFields(log.Fields{
				"secret":    secrets[0].Name,
				"namespace": secrets[0].Namespace,
			}).Info("Trident secret found by label.")
		} else {
			log.Warnf("Multiple secrets found matching label: %s; removing all.", appLabel)
		}

		if err = i.RemoveMultipleSecrets(secrets); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) RemoveMultipleSecrets(unwantedSecrets []v1.Secret) error {
	var err error
	var anyError bool
	var undeletedSecrets []string

	if len(unwantedSecrets) > 0 {
		for _, secretToRemove := range unwantedSecrets {
			// Delete the secret
			if err = i.client.DeleteSecret(secretToRemove.Name, secretToRemove.Namespace); err != nil {
				log.WithFields(log.Fields{
					"service":   secretToRemove.Name,
					"namespace": secretToRemove.Namespace,
					"label":     appLabel,
					"error":     err,
				}).Warning("Could not delete Trident secret.")

				anyError = true
				undeletedSecrets = append(undeletedSecrets, fmt.Sprintf("%v/%v", secretToRemove.Namespace,
					secretToRemove.Name))
			} else {
				log.Infof("Deleted Trident secret: %v/%v", secretToRemove.Namespace, secretToRemove.Name)
			}
		}
	}

	if anyError {
		return fmt.Errorf("unable to delete Trident secret(s): %v", undeletedSecrets)
	}

	return nil
}

func (i *Installer) deletePodSecurityPolicy() error {

	if podSecurityPolicies, err := i.client.GetPodSecurityPoliciesByLabel(appLabel); err != nil {
		log.Errorf("Unable to get list of Pod security policies by label %v", appLabel)
		return fmt.Errorf("unable to get list of Pod security policies")
	} else if len(podSecurityPolicies) == 0 {

		log.WithFields(log.Fields{
			"label": appLabel,
			"error": err,
		}).Warning("Trident Pod security policy not found.")
	} else {
		if len(podSecurityPolicies) == 1 {
			log.WithFields(log.Fields{
				"service":   podSecurityPolicies[0].Name,
				"namespace": podSecurityPolicies[0].Namespace,
			}).Info("Trident Pod security policy found by label.")
		} else {
			log.Warnf("Multiple Pod security policies found matching label: %s; removing all.", appLabel)
		}

		if err = i.RemoveMultiplePodSecurityPolicies(podSecurityPolicies); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) RemoveMultiplePodSecurityPolicies(unwantedPSPs []v1beta1.PodSecurityPolicy) error {
	var err error
	var anyError bool
	var undeletedPSPs []string

	if len(unwantedPSPs) > 0 {
		// Delete the pod security policies
		for _, PSPsToRemove := range unwantedPSPs {
			if err = i.client.DeletePodSecurityPolicy(PSPsToRemove.Name); err != nil {
				log.WithFields(log.Fields{
					"podSecurityPolicy": PSPsToRemove.Name,
					"label":             appLabel,
					"error":             err,
				}).Warning("Could not delete Trident Pod security policy.")

				anyError = true
				undeletedPSPs = append(undeletedPSPs, fmt.Sprintf("%v", PSPsToRemove.Name))
			} else {
				log.Infof("Deleted Trident Pod security policy: %v", PSPsToRemove.Name)
			}
		}
	}

	if anyError {
		return fmt.Errorf("unable to delete Trident service(s): %v", undeletedPSPs)
	}

	return nil
}

func (i *Installer) deleteTridentCSIDriverCR() error {

	if csiDrivers, err := i.client.GetCSIDriversByLabel(appLabel); err != nil {
		log.Errorf("Unable to get list of CSI driver CRs by label %v", appLabel)
		return fmt.Errorf("unable to get list of CSI driver CRs by label")
	} else if len(csiDrivers) == 0 {
		log.WithFields(log.Fields{
			"label": appLabel,
			"error": err,
		}).Warning("CSI driver CR not found.")
	} else {
		if len(csiDrivers) == 1 {
			log.WithFields(log.Fields{
				"service":   csiDrivers[0].Name,
				"namespace": csiDrivers[0].Namespace,
			}).Info("Trident CSI driver CR found by label.")
		} else {
			log.Warnf("Multiple CSI driver CRs found matching label: %s; removing all.", appLabel)
		}

		if err = i.RemoveMultipleCSIDriverCRs(csiDrivers); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) RemoveMultipleCSIDriverCRs(unwantedCSIDriverCRs []v1beta12.CSIDriver) error {
	var err error
	var anyError bool
	var undeletedCSIDriverCRs []string

	if len(unwantedCSIDriverCRs) > 0 {
		// Delete the CSI driver CRs
		for _, CSIDriverCRToRemove := range unwantedCSIDriverCRs {
			if err = i.client.DeleteCSIDriver(CSIDriverCRToRemove.Name); err != nil {
				log.WithFields(log.Fields{
					"CSIDriver": CSIDriverCRToRemove.Name,
					"label":     appLabel,
					"error":     err,
				}).Warning("Could not delete CSI driver CR.")

				anyError = true
				undeletedCSIDriverCRs = append(undeletedCSIDriverCRs, fmt.Sprintf("%v", CSIDriverCRToRemove.Name))
			} else {
				log.Infof("Deleted CSI driver CR: %v", CSIDriverCRToRemove.Name)
			}
		}
	}

	if anyError {
		return fmt.Errorf("unable to delete CSI driver CR(s): %v", undeletedCSIDriverCRs)
	}

	return nil
}

func (i *Installer) deleteTridentClusterRoleBinding() error {

	// Delete cluster role binding
	if clusterRoleBindings, err := i.client.GetClusterRoleBindingsByLabel(appLabel); err != nil {
		log.Errorf("Unable to get list of Pod security policies by label %v", appLabel)
		return fmt.Errorf("unable to get list of Pod security policies")
	} else if len(clusterRoleBindings) == 0 {
		log.WithFields(log.Fields{
			"label": appLabel,
			"error": err,
		}).Warning("Trident cluster role binding not found.")
	} else {
		if len(clusterRoleBindings) == 1 {
			log.WithFields(log.Fields{
				"clusterRoleBinding": clusterRoleBindings[0].Name,
				"namespace":          clusterRoleBindings[0].Namespace,
			}).Info("Trident Cluster role binding found by label.")
		} else {
			log.Warnf("Multiple Cluster role bindings found matching label: %s; removing all.", appLabel)
		}

		if err = i.RemoveMultipleClusterRoleBindings(clusterRoleBindings); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) RemoveMultipleClusterRoleBindings(unwantedClusterRoleBindings []v12.ClusterRoleBinding) error {
	var err error
	var anyError bool
	var undeletedClusterRoleBindings []string

	if len(unwantedClusterRoleBindings) > 0 {
		// Delete the cluster roles bindings
		for _, clusterRoleBindingToRemove := range unwantedClusterRoleBindings {
			if err = i.client.DeleteClusterRoleBinding(clusterRoleBindingToRemove.Name); err != nil {
				log.WithFields(log.Fields{
					"clusterRoleBinding": clusterRoleBindingToRemove.Name,
					"label":              appLabel,
					"error":              err,
				}).Warning("Could not delete Trident cluster role binding.")

				anyError = true
				undeletedClusterRoleBindings = append(undeletedClusterRoleBindings, fmt.Sprintf("%v", clusterRoleBindingToRemove.Name))
			} else {
				log.Infof("Deleted Trident Cluster role binding: %v", clusterRoleBindingToRemove.Name)
			}
		}
	}

	if anyError {
		return fmt.Errorf("unable to delete Trident cluster role binding(s): %v", undeletedClusterRoleBindings)
	}

	return nil
}

func (i *Installer) deleteTridentClusterRole() error {

	// Delete cluster role
	if clusterRoles, err := i.client.GetClusterRolesByLabel(appLabel); err != nil {
		log.Errorf("Unable to get list of Cluster roles by label %v", appLabel)
		return fmt.Errorf("unable to get list of Cluster roles")
	} else if len(clusterRoles) == 0 {
		log.WithFields(log.Fields{
			"label": appLabel,
			"error": err,
		}).Warning("Trident cluster role not found.")
	} else {
		if len(clusterRoles) == 1 {
			log.WithFields(log.Fields{
				"clusterRole": clusterRoles[0].Name,
				"namespace":   clusterRoles[0].Namespace,
			}).Info("Trident Cluster role found by label.")
		} else {
			log.Warnf("Multiple Cluster roles found matching label: %s; removing all.", appLabel)
		}

		if err = i.RemoveMultipleClusterRoles(clusterRoles); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) RemoveMultipleClusterRoles(unwantedClusterRoles []v12.ClusterRole) error {
	var err error
	var anyError bool
	var undeletedClusterRoles []string

	if len(unwantedClusterRoles) > 0 {
		// Delete the cluster roles
		for _, clusterRoleToRemove := range unwantedClusterRoles {
			if err = i.client.DeleteClusterRole(clusterRoleToRemove.Name); err != nil {
				log.WithFields(log.Fields{
					"clusterRole": clusterRoleToRemove.Name,
					"label":       appLabel,
					"error":       err,
				}).Warning("Could not delete Trident cluster role.")

				anyError = true
				undeletedClusterRoles = append(undeletedClusterRoles, fmt.Sprintf("%v", clusterRoleToRemove.Name))
			} else {
				log.Infof("Deleted Trident Cluster role: %v", clusterRoleToRemove.Name)
			}
		}
	}

	if anyError {
		return fmt.Errorf("unable to delete Trident cluster role(s): %v", undeletedClusterRoles)
	}

	return nil
}

func (i *Installer) deleteTridentServiceAccount() error {

	// Delete service account
	if serviceAccounts, err := i.client.GetServiceAccountsByLabel(appLabel, false); err != nil {
		log.Errorf("Unable to get list of Service accounts by label %v", appLabel)
		return fmt.Errorf("unable to get list of Service accounts")
	} else if len(serviceAccounts) == 0 {
		log.WithFields(log.Fields{
			"label": appLabel,
			"error": err,
		}).Warning("Trident service account not found.")
	} else {
		if len(serviceAccounts) == 1 {
			log.WithFields(log.Fields{
				"service":   serviceAccounts[0].Name,
				"namespace": serviceAccounts[0].Namespace,
			}).Info("Trident Service accounts found by label.")
		} else {
			log.Warnf("Multiple Service accounts found matching label: %s; removing all.", appLabel)
		}

		if err = i.RemoveMultipleServiceAccounts(serviceAccounts); err != nil {
			return err
		}
	}

	return nil
}

func (i *Installer) RemoveMultipleServiceAccounts(unwantedServiceAccounts []v1.ServiceAccount) error {
	var err error
	var anyError bool
	var undeletedServiceAccounts []string

	if len(unwantedServiceAccounts) > 0 {
		// Delete the service accounts
		for _, serviceAccountToRemove := range unwantedServiceAccounts {
			if err = i.client.DeleteServiceAccount(serviceAccountToRemove.Name,
				serviceAccountToRemove.Namespace); err != nil {
				log.WithFields(log.Fields{
					"serviceAccount": serviceAccountToRemove.Name,
					"namespace":      serviceAccountToRemove.Namespace,
					"label":          appLabel,
					"error":          err,
				}).Errorf("Could not delete Trident service account.")

				anyError = true
				undeletedServiceAccounts = append(undeletedServiceAccounts, fmt.Sprintf("%v/%v", serviceAccountToRemove.Namespace,
					serviceAccountToRemove.Name))
			} else {
				log.Infof("Deleted Trident service account: %v/%v", serviceAccountToRemove.Namespace,
					serviceAccountToRemove.Name)
			}
		}
	}

	if anyError {
		return fmt.Errorf("unable to delete Trident service account(s): %v", undeletedServiceAccounts)
	}

	return nil
}

func (i *Installer) deleteCustomResourceDefinitions() (returnError error) {

	var logFields log.Fields

	returnError = i.client.DeleteObjectByYAML(k8sclient.GetCRDsYAML(false), false)
	if returnError != nil {
		returnError = fmt.Errorf("could not delete custom resource definitions in %s; %v", i.namespace, returnError)
		return
	}
	log.WithFields(logFields).Info("Deleted custom resource definitions.")
	return nil
}

func (i *Installer) deleteCustomResourceDefinition(crdName, crdYAML string) (returnError error) {

	var logFields log.Fields

	returnError = i.client.DeleteObjectByYAML(crdYAML, false)
	if returnError != nil {
		returnError = fmt.Errorf("could not delete custom resource definition %v in %s; %v", crdName, i.namespace,
			returnError)
		return
	}
	log.WithFields(logFields).Infof("Deleted custom resource definitions %v.", crdName)
	return nil
}

func (i *Installer) RemoveMultiplePods(unwantedPods []v1.Pod) error {
	var err error
	var anyError bool
	var undeletedPods []string

	if len(unwantedPods) > 0 {
		for _, podToRemove := range unwantedPods {
			// Delete the pod
			if err = i.client.DeletePod(podToRemove.Name, podToRemove.Namespace); err != nil {
				log.WithFields(log.Fields{
					"service":   podToRemove.Name,
					"namespace": podToRemove.Namespace,
					"error":     err,
				}).Warning("Could not delete pod.")

				anyError = true
				undeletedPods = append(undeletedPods, fmt.Sprintf("%v/%v", podToRemove.Namespace,
					podToRemove.Name))
			} else {
				log.Infof("Deleted pod: %v/%v", podToRemove.Namespace, podToRemove.Name)
			}
		}
	}

	if anyError {
		return fmt.Errorf("unable to delete pod(s): %v", undeletedPods)
	}

	return nil
}

func (i *Installer) ObliviateCRDs() error {
	return cmd.ObliviateCRDs(i.client, i.tridentCRDClient, i.namespace, k8sTimeout)
}
