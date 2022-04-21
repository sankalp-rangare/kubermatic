// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NUMA n u m a
//
// swagger:model NUMA
type NUMA struct {

	// guest mapping passthrough
	GuestMappingPassthrough NUMAGuestMappingPassthrough `json:"guestMappingPassthrough,omitempty"`
}

// Validate validates this n u m a
func (m *NUMA) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this n u m a based on context it is used
func (m *NUMA) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NUMA) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NUMA) UnmarshalBinary(b []byte) error {
	var res NUMA
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}