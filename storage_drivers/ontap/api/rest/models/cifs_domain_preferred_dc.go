// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CifsDomainPreferredDc cifs domain preferred dc
//
// swagger:model cifs_domain_preferred_dc
type CifsDomainPreferredDc struct {

	// Fully Qualified Domain Name.
	//
	// Example: test.com
	Fqdn string `json:"fqdn,omitempty"`

	// IP address of the preferred domain controller (DC). The address can be either an IPv4 or an IPv6 address.
	//
	// Example: 4.4.4.4
	ServerIP string `json:"server_ip,omitempty"`
}

// Validate validates this cifs domain preferred dc
func (m *CifsDomainPreferredDc) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cifs domain preferred dc based on context it is used
func (m *CifsDomainPreferredDc) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CifsDomainPreferredDc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CifsDomainPreferredDc) UnmarshalBinary(b []byte) error {
	var res CifsDomainPreferredDc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}