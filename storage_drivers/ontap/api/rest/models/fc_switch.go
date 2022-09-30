// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FcSwitch A Fibre Channel switch.
//
// swagger:model fc_switch
type FcSwitch struct {

	// links
	Links *FcSwitchLinks `json:"_links,omitempty"`

	// cache
	Cache *FcSwitchCache `json:"cache,omitempty"`

	// The domain identifier (ID) of the Fibre Channel (FC) switch. The domain ID is a unique identifier for the FC switch in the FC fabric.
	//
	// Example: 1
	// Read Only: true
	// Maximum: 239
	// Minimum: 1
	DomainID int64 `json:"domain_id,omitempty"`

	// fabric
	Fabric *FcSwitchFabric `json:"fabric,omitempty"`

	// The logical name of the Fibre Channel switch.
	//
	// Example: switch1
	// Read Only: true
	Name string `json:"name,omitempty"`

	// An array of the Fibre Channel (FC) switch's ports and their attached FC devices.
	//
	Ports []*FcSwitchPortsItems0 `json:"ports,omitempty"`

	// The firmware release of the Fibre Channel switch.
	//
	// Example: 1.0.
	// Read Only: true
	Release string `json:"release,omitempty"`

	// The vendor of the Fibre Channel switch.
	//
	// Example: vendor1
	// Read Only: true
	Vendor string `json:"vendor,omitempty"`

	// The world-wide name (WWN) for the Fibre Channel switch.
	//
	// Example: 10:00:e1:e2:e3:e4:e5:e6
	// Read Only: true
	Wwn string `json:"wwn,omitempty"`
}

// Validate validates this fc switch
func (m *FcSwitch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomainID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFabric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitch) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitch) validateCache(formats strfmt.Registry) error {
	if swag.IsZero(m.Cache) { // not required
		return nil
	}

	if m.Cache != nil {
		if err := m.Cache.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cache")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitch) validateDomainID(formats strfmt.Registry) error {
	if swag.IsZero(m.DomainID) { // not required
		return nil
	}

	if err := validate.MinimumInt("domain_id", "body", m.DomainID, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("domain_id", "body", m.DomainID, 239, false); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitch) validateFabric(formats strfmt.Registry) error {
	if swag.IsZero(m.Fabric) { // not required
		return nil
	}

	if m.Fabric != nil {
		if err := m.Fabric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitch) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this fc switch based on the context it is used
func (m *FcSwitch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCache(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDomainID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFabric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelease(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVendor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitch) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitch) contextValidateCache(ctx context.Context, formats strfmt.Registry) error {

	if m.Cache != nil {
		if err := m.Cache.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cache")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitch) contextValidateDomainID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "domain_id", "body", int64(m.DomainID)); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitch) contextValidateFabric(ctx context.Context, formats strfmt.Registry) error {

	if m.Fabric != nil {
		if err := m.Fabric.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitch) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitch) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if m.Ports[i] != nil {
			if err := m.Ports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FcSwitch) contextValidateRelease(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "release", "body", string(m.Release)); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitch) contextValidateVendor(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "vendor", "body", string(m.Vendor)); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitch) contextValidateWwn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wwn", "body", string(m.Wwn)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitch) UnmarshalBinary(b []byte) error {
	var res FcSwitch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchCache Properties of Fibre Chanel fabric cache.
//
// swagger:model FcSwitchCache
type FcSwitchCache struct {

	// The age of the Fibre Channel fabric data cache retrieved. If the FC fabric data cache has not been fully updated for a newly discovered fabric, or a fabric that has been re-discovered after being purged, a value for this property will not be retrieved. The value is in ISO 8601 duration format.
	//
	// Example: PT3M30S
	// Read Only: true
	Age string `json:"age,omitempty"`

	// A boolean that indicates if the retrieved data is current relative to the `cache.maximum_age` value of the request. A value of `true` indicates that the data is no older than the requested maximum age. A value of `false` indicates that the data is older than the requested maximum age; if more current data is required, the caller should wait for some time for the cache update to complete and query the data again.
	//
	// Read Only: true
	IsCurrent *bool `json:"is_current,omitempty"`

	// The date and time at which the Fibre Channel fabric data cache retrieved was last updated. If the FC fabric data cache has not been fully updated for a newly discovered fabric, or a fabric that has been re-discovered after being purged, a value for this property will not be retrieved.
	//
	// Read Only: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this fc switch cache
func (m *FcSwitchCache) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchCache) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("cache"+"."+"update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fc switch cache based on the context it is used
func (m *FcSwitchCache) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsCurrent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchCache) contextValidateAge(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cache"+"."+"age", "body", string(m.Age)); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchCache) contextValidateIsCurrent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cache"+"."+"is_current", "body", m.IsCurrent); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchCache) contextValidateUpdateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cache"+"."+"update_time", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchCache) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchCache) UnmarshalBinary(b []byte) error {
	var res FcSwitchCache
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchFabric A reference to a Fibre Channel fabric.
//
// swagger:model FcSwitchFabric
type FcSwitchFabric struct {

	// links
	Links *FcSwitchFabricLinks `json:"_links,omitempty"`

	// The world wide name (WWN) of the primary switch of the Fibre Channel (FC) fabric. This is used as a unique identifier for the FC fabric.
	//
	// Example: 10:00:d1:d2:d3:d4:d5:d6
	Name string `json:"name,omitempty"`
}

// Validate validates this fc switch fabric
func (m *FcSwitchFabric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchFabric) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc switch fabric based on the context it is used
func (m *FcSwitchFabric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchFabric) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchFabric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchFabric) UnmarshalBinary(b []byte) error {
	var res FcSwitchFabric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchFabricLinks fc switch fabric links
//
// swagger:model FcSwitchFabricLinks
type FcSwitchFabricLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fc switch fabric links
func (m *FcSwitchFabricLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchFabricLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc switch fabric links based on the context it is used
func (m *FcSwitchFabricLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchFabricLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fabric" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchFabricLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchFabricLinks) UnmarshalBinary(b []byte) error {
	var res FcSwitchFabricLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchLinks fc switch links
//
// swagger:model FcSwitchLinks
type FcSwitchLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fc switch links
func (m *FcSwitchLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc switch links based on the context it is used
func (m *FcSwitchLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchLinks) UnmarshalBinary(b []byte) error {
	var res FcSwitchLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchPortsItems0 A Fibre Channel switch port.
//
// swagger:model FcSwitchPortsItems0
type FcSwitchPortsItems0 struct {

	// attached device
	AttachedDevice *FcSwitchPortsItems0AttachedDevice `json:"attached_device,omitempty"`

	// The slot of the Fibre Channel switch port.
	//
	// Example: 1
	// Read Only: true
	Slot string `json:"slot,omitempty"`

	// The state of the Fibre Channel switch port.
	//
	// Example: online
	// Read Only: true
	// Enum: [unknown online offline testing fault]
	State string `json:"state,omitempty"`

	// The type of the Fibre Channel switch port.
	//
	// Read Only: true
	// Enum: [b_port e_port f_port fl_port fnl_port fv_port n_port nl_port nv_port nx_port sd_port te_port tf_port tl_port tnp_port none]
	Type string `json:"type,omitempty"`

	// The world wide port name (WWPN) of the Fibre Channel switch port.
	//
	// Example: 50:0a:31:32:33:34:35:36
	// Read Only: true
	Wwpn string `json:"wwpn,omitempty"`
}

// Validate validates this fc switch ports items0
func (m *FcSwitchPortsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachedDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *FcSwitchPortsItems0) validateAttachedDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.AttachedDevice) { // not required
		return nil
	}

	if m.AttachedDevice != nil {
		if err := m.AttachedDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attached_device")
			}
			return err
		}
	}

	return nil
}

var fcSwitchPortsItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","online","offline","testing","fault"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcSwitchPortsItems0TypeStatePropEnum = append(fcSwitchPortsItems0TypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// state
	// State
	// unknown
	// END DEBUGGING
	// FcSwitchPortsItems0StateUnknown captures enum value "unknown"
	FcSwitchPortsItems0StateUnknown string = "unknown"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// state
	// State
	// online
	// END DEBUGGING
	// FcSwitchPortsItems0StateOnline captures enum value "online"
	FcSwitchPortsItems0StateOnline string = "online"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// state
	// State
	// offline
	// END DEBUGGING
	// FcSwitchPortsItems0StateOffline captures enum value "offline"
	FcSwitchPortsItems0StateOffline string = "offline"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// state
	// State
	// testing
	// END DEBUGGING
	// FcSwitchPortsItems0StateTesting captures enum value "testing"
	FcSwitchPortsItems0StateTesting string = "testing"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// state
	// State
	// fault
	// END DEBUGGING
	// FcSwitchPortsItems0StateFault captures enum value "fault"
	FcSwitchPortsItems0StateFault string = "fault"
)

// prop value enum
func (m *FcSwitchPortsItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcSwitchPortsItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcSwitchPortsItems0) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var fcSwitchPortsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["b_port","e_port","f_port","fl_port","fnl_port","fv_port","n_port","nl_port","nv_port","nx_port","sd_port","te_port","tf_port","tl_port","tnp_port","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcSwitchPortsItems0TypeTypePropEnum = append(fcSwitchPortsItems0TypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// b_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeBPort captures enum value "b_port"
	FcSwitchPortsItems0TypeBPort string = "b_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// e_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeEPort captures enum value "e_port"
	FcSwitchPortsItems0TypeEPort string = "e_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// f_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeFPort captures enum value "f_port"
	FcSwitchPortsItems0TypeFPort string = "f_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// fl_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeFlPort captures enum value "fl_port"
	FcSwitchPortsItems0TypeFlPort string = "fl_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// fnl_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeFnlPort captures enum value "fnl_port"
	FcSwitchPortsItems0TypeFnlPort string = "fnl_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// fv_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeFvPort captures enum value "fv_port"
	FcSwitchPortsItems0TypeFvPort string = "fv_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// n_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeNPort captures enum value "n_port"
	FcSwitchPortsItems0TypeNPort string = "n_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// nl_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeNlPort captures enum value "nl_port"
	FcSwitchPortsItems0TypeNlPort string = "nl_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// nv_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeNvPort captures enum value "nv_port"
	FcSwitchPortsItems0TypeNvPort string = "nv_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// nx_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeNxPort captures enum value "nx_port"
	FcSwitchPortsItems0TypeNxPort string = "nx_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// sd_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeSdPort captures enum value "sd_port"
	FcSwitchPortsItems0TypeSdPort string = "sd_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// te_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeTePort captures enum value "te_port"
	FcSwitchPortsItems0TypeTePort string = "te_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// tf_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeTfPort captures enum value "tf_port"
	FcSwitchPortsItems0TypeTfPort string = "tf_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// tl_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeTlPort captures enum value "tl_port"
	FcSwitchPortsItems0TypeTlPort string = "tl_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// tnp_port
	// END DEBUGGING
	// FcSwitchPortsItems0TypeTnpPort captures enum value "tnp_port"
	FcSwitchPortsItems0TypeTnpPort string = "tnp_port"

	// BEGIN DEBUGGING
	// FcSwitchPortsItems0
	// FcSwitchPortsItems0
	// type
	// Type
	// none
	// END DEBUGGING
	// FcSwitchPortsItems0TypeNone captures enum value "none"
	FcSwitchPortsItems0TypeNone string = "none"
)

// prop value enum
func (m *FcSwitchPortsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcSwitchPortsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcSwitchPortsItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fc switch ports items0 based on the context it is used
func (m *FcSwitchPortsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachedDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchPortsItems0) contextValidateAttachedDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.AttachedDevice != nil {
		if err := m.AttachedDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attached_device")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitchPortsItems0) contextValidateSlot(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "slot", "body", string(m.Slot)); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchPortsItems0) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchPortsItems0) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchPortsItems0) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wwpn", "body", string(m.Wwpn)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchPortsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchPortsItems0) UnmarshalBinary(b []byte) error {
	var res FcSwitchPortsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchPortsItems0AttachedDevice The Fibre Channel (FC) device attached to the FC switch port.
//
// swagger:model FcSwitchPortsItems0AttachedDevice
type FcSwitchPortsItems0AttachedDevice struct {

	// The Fibre Channel port identifier of the attach device.
	//
	// Example: 0x011300
	// Read Only: true
	PortID string `json:"port_id,omitempty"`

	// The world-wide port name (WWPN) of the attached device.
	//
	// Example: 50:0a:21:22:23:24:25:26
	// Read Only: true
	Wwpn string `json:"wwpn,omitempty"`
}

// Validate validates this fc switch ports items0 attached device
func (m *FcSwitchPortsItems0AttachedDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fc switch ports items0 attached device based on the context it is used
func (m *FcSwitchPortsItems0AttachedDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePortID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchPortsItems0AttachedDevice) contextValidatePortID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "attached_device"+"."+"port_id", "body", string(m.PortID)); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchPortsItems0AttachedDevice) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "attached_device"+"."+"wwpn", "body", string(m.Wwpn)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchPortsItems0AttachedDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchPortsItems0AttachedDevice) UnmarshalBinary(b []byte) error {
	var res FcSwitchPortsItems0AttachedDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}