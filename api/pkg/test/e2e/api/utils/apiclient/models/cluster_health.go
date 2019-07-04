// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ClusterHealth ClusterHealth stores health information about the cluster's components.
// swagger:model ClusterHealth
type ClusterHealth struct {

	// apiserver
	Apiserver bool `json:"apiserver,omitempty"`

	// cloud provider infrastructure
	CloudProviderInfrastructure bool `json:"cloudProviderInfrastructure,omitempty"`

	// controller
	Controller bool `json:"controller,omitempty"`

	// etcd
	Etcd bool `json:"etcd,omitempty"`

	// machine controller
	MachineController bool `json:"machineController,omitempty"`

	// scheduler
	Scheduler bool `json:"scheduler,omitempty"`

	// user cluster controller manager
	UserClusterControllerManager bool `json:"userClusterControllerManager,omitempty"`
}

// Validate validates this cluster health
func (m *ClusterHealth) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterHealth) UnmarshalBinary(b []byte) error {
	var res ClusterHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}