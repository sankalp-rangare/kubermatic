// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Hugepages Hugepages allow to use hugepages for the VirtualMachineInstance instead of regular memory.
//
// swagger:model Hugepages
type Hugepages struct {

	// PageSize specifies the hugepage size, for x86_64 architecture valid values are 1Gi and 2Mi.
	PageSize string `json:"pageSize,omitempty"`
}

// Validate validates this hugepages
func (m *Hugepages) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hugepages based on context it is used
func (m *Hugepages) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Hugepages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hugepages) UnmarshalBinary(b []byte) error {
	var res Hugepages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}