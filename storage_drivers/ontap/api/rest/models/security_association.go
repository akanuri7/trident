// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SecurityAssociation Security association object for IPsec security association and IKE (Internet Key Exchange) security association.
//
// swagger:model security_association
type SecurityAssociation struct {

	// Cipher suite for the security association.
	// Enum: [suite_aescbc suiteb_gcm256 suiteb_gmac256]
	CipherSuite *string `json:"cipher_suite,omitempty"`

	// ike
	Ike *SecurityAssociationIke `json:"ike,omitempty"`

	// ipsec
	Ipsec *SecurityAssociationIpsec `json:"ipsec,omitempty"`

	// Lifetime for the security association in seconds.
	Lifetime int64 `json:"lifetime,omitempty"`

	// Local address of the security association.
	LocalAddress string `json:"local_address,omitempty"`

	// node
	Node *SecurityAssociationNode `json:"node,omitempty"`

	// Policy name for the security association.
	PolicyName string `json:"policy_name,omitempty"`

	// Remote address of the security association.
	RemoteAddress string `json:"remote_address,omitempty"`

	// scope
	Scope NetworkScopeReadonly `json:"scope,omitempty"`

	// svm
	Svm *SecurityAssociationSvm `json:"svm,omitempty"`

	// Type of security association, it can be IPsec or IKE (Internet Key Exchange).
	// Enum: [ipsec ike]
	Type *string `json:"type,omitempty"`

	// Unique identifier of the security association.
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this security association
func (m *SecurityAssociation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCipherSuite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIke(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpsec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var securityAssociationTypeCipherSuitePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["suite_aescbc","suiteb_gcm256","suiteb_gmac256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityAssociationTypeCipherSuitePropEnum = append(securityAssociationTypeCipherSuitePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// security_association
	// SecurityAssociation
	// cipher_suite
	// CipherSuite
	// suite_aescbc
	// END DEBUGGING
	// SecurityAssociationCipherSuiteSuiteAescbc captures enum value "suite_aescbc"
	SecurityAssociationCipherSuiteSuiteAescbc string = "suite_aescbc"

	// BEGIN DEBUGGING
	// security_association
	// SecurityAssociation
	// cipher_suite
	// CipherSuite
	// suiteb_gcm256
	// END DEBUGGING
	// SecurityAssociationCipherSuiteSuitebGcm256 captures enum value "suiteb_gcm256"
	SecurityAssociationCipherSuiteSuitebGcm256 string = "suiteb_gcm256"

	// BEGIN DEBUGGING
	// security_association
	// SecurityAssociation
	// cipher_suite
	// CipherSuite
	// suiteb_gmac256
	// END DEBUGGING
	// SecurityAssociationCipherSuiteSuitebGmac256 captures enum value "suiteb_gmac256"
	SecurityAssociationCipherSuiteSuitebGmac256 string = "suiteb_gmac256"
)

// prop value enum
func (m *SecurityAssociation) validateCipherSuiteEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityAssociationTypeCipherSuitePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityAssociation) validateCipherSuite(formats strfmt.Registry) error {
	if swag.IsZero(m.CipherSuite) { // not required
		return nil
	}

	// value enum
	if err := m.validateCipherSuiteEnum("cipher_suite", "body", *m.CipherSuite); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAssociation) validateIke(formats strfmt.Registry) error {
	if swag.IsZero(m.Ike) { // not required
		return nil
	}

	if m.Ike != nil {
		if err := m.Ike.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ike")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityAssociation) validateIpsec(formats strfmt.Registry) error {
	if swag.IsZero(m.Ipsec) { // not required
		return nil
	}

	if m.Ipsec != nil {
		if err := m.Ipsec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipsec")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityAssociation) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityAssociation) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if err := m.Scope.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		}
		return err
	}

	return nil
}

func (m *SecurityAssociation) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

var securityAssociationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipsec","ike"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityAssociationTypeTypePropEnum = append(securityAssociationTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// security_association
	// SecurityAssociation
	// type
	// Type
	// ipsec
	// END DEBUGGING
	// SecurityAssociationTypeIpsec captures enum value "ipsec"
	SecurityAssociationTypeIpsec string = "ipsec"

	// BEGIN DEBUGGING
	// security_association
	// SecurityAssociation
	// type
	// Type
	// ike
	// END DEBUGGING
	// SecurityAssociationTypeIke captures enum value "ike"
	SecurityAssociationTypeIke string = "ike"
)

// prop value enum
func (m *SecurityAssociation) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityAssociationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityAssociation) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this security association based on the context it is used
func (m *SecurityAssociation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIke(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIpsec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAssociation) contextValidateIke(ctx context.Context, formats strfmt.Registry) error {

	if m.Ike != nil {
		if err := m.Ike.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ike")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityAssociation) contextValidateIpsec(ctx context.Context, formats strfmt.Registry) error {

	if m.Ipsec != nil {
		if err := m.Ipsec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipsec")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityAssociation) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityAssociation) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Scope.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		}
		return err
	}

	return nil
}

func (m *SecurityAssociation) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAssociation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAssociation) UnmarshalBinary(b []byte) error {
	var res SecurityAssociation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAssociationIke Objects containing parameters specific to IKE (Internet Key Exchange) security association.
//
// swagger:model SecurityAssociationIke
type SecurityAssociationIke struct {

	// Authentication method for internet key exchange protocol.
	// Enum: [none psk cert]
	Authentication *string `json:"authentication,omitempty"`

	// Initiator's security parameter index for the IKE security association.
	InitiatorSecurityParameterIndex string `json:"initiator_security_parameter_index,omitempty"`

	// Indicates whether or not IKE has been initiated by this node.
	IsInitiator bool `json:"is_initiator,omitempty"`

	// Responder's security parameter index for the IKE security association.
	ResponderSecurityParameterIndex string `json:"responder_security_parameter_index,omitempty"`

	// State of the IKE connection.
	// Enum: [none connecting established dead_peer_probe]
	State string `json:"state,omitempty"`

	// Internet key exchange protocol version.
	Version int64 `json:"version,omitempty"`
}

// Validate validates this security association ike
func (m *SecurityAssociationIke) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var securityAssociationIkeTypeAuthenticationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","psk","cert"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityAssociationIkeTypeAuthenticationPropEnum = append(securityAssociationIkeTypeAuthenticationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// SecurityAssociationIke
	// SecurityAssociationIke
	// authentication
	// Authentication
	// none
	// END DEBUGGING
	// SecurityAssociationIkeAuthenticationNone captures enum value "none"
	SecurityAssociationIkeAuthenticationNone string = "none"

	// BEGIN DEBUGGING
	// SecurityAssociationIke
	// SecurityAssociationIke
	// authentication
	// Authentication
	// psk
	// END DEBUGGING
	// SecurityAssociationIkeAuthenticationPsk captures enum value "psk"
	SecurityAssociationIkeAuthenticationPsk string = "psk"

	// BEGIN DEBUGGING
	// SecurityAssociationIke
	// SecurityAssociationIke
	// authentication
	// Authentication
	// cert
	// END DEBUGGING
	// SecurityAssociationIkeAuthenticationCert captures enum value "cert"
	SecurityAssociationIkeAuthenticationCert string = "cert"
)

// prop value enum
func (m *SecurityAssociationIke) validateAuthenticationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityAssociationIkeTypeAuthenticationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityAssociationIke) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationEnum("ike"+"."+"authentication", "body", *m.Authentication); err != nil {
		return err
	}

	return nil
}

var securityAssociationIkeTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","connecting","established","dead_peer_probe"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityAssociationIkeTypeStatePropEnum = append(securityAssociationIkeTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// SecurityAssociationIke
	// SecurityAssociationIke
	// state
	// State
	// none
	// END DEBUGGING
	// SecurityAssociationIkeStateNone captures enum value "none"
	SecurityAssociationIkeStateNone string = "none"

	// BEGIN DEBUGGING
	// SecurityAssociationIke
	// SecurityAssociationIke
	// state
	// State
	// connecting
	// END DEBUGGING
	// SecurityAssociationIkeStateConnecting captures enum value "connecting"
	SecurityAssociationIkeStateConnecting string = "connecting"

	// BEGIN DEBUGGING
	// SecurityAssociationIke
	// SecurityAssociationIke
	// state
	// State
	// established
	// END DEBUGGING
	// SecurityAssociationIkeStateEstablished captures enum value "established"
	SecurityAssociationIkeStateEstablished string = "established"

	// BEGIN DEBUGGING
	// SecurityAssociationIke
	// SecurityAssociationIke
	// state
	// State
	// dead_peer_probe
	// END DEBUGGING
	// SecurityAssociationIkeStateDeadPeerProbe captures enum value "dead_peer_probe"
	SecurityAssociationIkeStateDeadPeerProbe string = "dead_peer_probe"
)

// prop value enum
func (m *SecurityAssociationIke) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityAssociationIkeTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityAssociationIke) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("ike"+"."+"state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this security association ike based on context it is used
func (m *SecurityAssociationIke) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAssociationIke) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAssociationIke) UnmarshalBinary(b []byte) error {
	var res SecurityAssociationIke
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAssociationIpsec Objects containing parameters specific to IPsec security association.
//
// swagger:model SecurityAssociationIpsec
type SecurityAssociationIpsec struct {

	// Action for the IPsec security association.
	// Enum: [bypass discard esp_transport]
	Action *string `json:"action,omitempty"`

	// inbound
	Inbound *SecurityAssociationIpsecInbound `json:"inbound,omitempty"`

	// outbound
	Outbound *SecurityAssociationIpsecOutbound `json:"outbound,omitempty"`

	// State of the IPsec security association.
	State string `json:"state,omitempty"`
}

// Validate validates this security association ipsec
func (m *SecurityAssociationIpsec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInbound(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutbound(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var securityAssociationIpsecTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bypass","discard","esp_transport"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityAssociationIpsecTypeActionPropEnum = append(securityAssociationIpsecTypeActionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// SecurityAssociationIpsec
	// SecurityAssociationIpsec
	// action
	// Action
	// bypass
	// END DEBUGGING
	// SecurityAssociationIpsecActionBypass captures enum value "bypass"
	SecurityAssociationIpsecActionBypass string = "bypass"

	// BEGIN DEBUGGING
	// SecurityAssociationIpsec
	// SecurityAssociationIpsec
	// action
	// Action
	// discard
	// END DEBUGGING
	// SecurityAssociationIpsecActionDiscard captures enum value "discard"
	SecurityAssociationIpsecActionDiscard string = "discard"

	// BEGIN DEBUGGING
	// SecurityAssociationIpsec
	// SecurityAssociationIpsec
	// action
	// Action
	// esp_transport
	// END DEBUGGING
	// SecurityAssociationIpsecActionEspTransport captures enum value "esp_transport"
	SecurityAssociationIpsecActionEspTransport string = "esp_transport"
)

// prop value enum
func (m *SecurityAssociationIpsec) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityAssociationIpsecTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityAssociationIpsec) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("ipsec"+"."+"action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAssociationIpsec) validateInbound(formats strfmt.Registry) error {
	if swag.IsZero(m.Inbound) { // not required
		return nil
	}

	if m.Inbound != nil {
		if err := m.Inbound.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipsec" + "." + "inbound")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityAssociationIpsec) validateOutbound(formats strfmt.Registry) error {
	if swag.IsZero(m.Outbound) { // not required
		return nil
	}

	if m.Outbound != nil {
		if err := m.Outbound.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipsec" + "." + "outbound")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security association ipsec based on the context it is used
func (m *SecurityAssociationIpsec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInbound(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutbound(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAssociationIpsec) contextValidateInbound(ctx context.Context, formats strfmt.Registry) error {

	if m.Inbound != nil {
		if err := m.Inbound.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipsec" + "." + "inbound")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityAssociationIpsec) contextValidateOutbound(ctx context.Context, formats strfmt.Registry) error {

	if m.Outbound != nil {
		if err := m.Outbound.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipsec" + "." + "outbound")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAssociationIpsec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAssociationIpsec) UnmarshalBinary(b []byte) error {
	var res SecurityAssociationIpsec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAssociationIpsecInbound Status for inbound parameters for the IPsec security association.
//
// swagger:model SecurityAssociationIpsecInbound
type SecurityAssociationIpsecInbound struct {

	// Number of inbound bytes for the IPsec security association.
	Bytes int64 `json:"bytes,omitempty"`

	// Number of inbound packets for the IPsec security association.
	Packets int64 `json:"packets,omitempty"`

	// Inbound security parameter index for the IPSec security association.
	SecurityParameterIndex string `json:"security_parameter_index,omitempty"`
}

// Validate validates this security association ipsec inbound
func (m *SecurityAssociationIpsecInbound) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this security association ipsec inbound based on context it is used
func (m *SecurityAssociationIpsecInbound) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAssociationIpsecInbound) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAssociationIpsecInbound) UnmarshalBinary(b []byte) error {
	var res SecurityAssociationIpsecInbound
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAssociationIpsecOutbound Status for outbound parameters for the IPsec security association.
//
// swagger:model SecurityAssociationIpsecOutbound
type SecurityAssociationIpsecOutbound struct {

	// Number of outbound bytes for the IPsec security association.
	Bytes int64 `json:"bytes,omitempty"`

	// Number of outbound packets for the IPsec security association.
	Packets int64 `json:"packets,omitempty"`

	// Outbound security parameter index for the IPSec security association.
	SecurityParameterIndex string `json:"security_parameter_index,omitempty"`
}

// Validate validates this security association ipsec outbound
func (m *SecurityAssociationIpsecOutbound) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this security association ipsec outbound based on context it is used
func (m *SecurityAssociationIpsecOutbound) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAssociationIpsecOutbound) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAssociationIpsecOutbound) UnmarshalBinary(b []byte) error {
	var res SecurityAssociationIpsecOutbound
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAssociationNode Node with the security association.
//
// swagger:model SecurityAssociationNode
type SecurityAssociationNode struct {

	// links
	Links *SecurityAssociationNodeLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this security association node
func (m *SecurityAssociationNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAssociationNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security association node based on the context it is used
func (m *SecurityAssociationNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAssociationNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAssociationNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAssociationNode) UnmarshalBinary(b []byte) error {
	var res SecurityAssociationNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAssociationNodeLinks security association node links
//
// swagger:model SecurityAssociationNodeLinks
type SecurityAssociationNodeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this security association node links
func (m *SecurityAssociationNodeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAssociationNodeLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security association node links based on the context it is used
func (m *SecurityAssociationNodeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAssociationNodeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAssociationNodeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAssociationNodeLinks) UnmarshalBinary(b []byte) error {
	var res SecurityAssociationNodeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAssociationSvm security association svm
//
// swagger:model SecurityAssociationSvm
type SecurityAssociationSvm struct {

	// links
	Links *SecurityAssociationSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this security association svm
func (m *SecurityAssociationSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAssociationSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security association svm based on the context it is used
func (m *SecurityAssociationSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAssociationSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAssociationSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAssociationSvm) UnmarshalBinary(b []byte) error {
	var res SecurityAssociationSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAssociationSvmLinks security association svm links
//
// swagger:model SecurityAssociationSvmLinks
type SecurityAssociationSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this security association svm links
func (m *SecurityAssociationSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAssociationSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security association svm links based on the context it is used
func (m *SecurityAssociationSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAssociationSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAssociationSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAssociationSvmLinks) UnmarshalBinary(b []byte) error {
	var res SecurityAssociationSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}