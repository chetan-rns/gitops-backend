// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1alpha1ApplicationStatus ApplicationStatus contains information about application sync, health status
//
// swagger:model v1alpha1ApplicationStatus
type V1alpha1ApplicationStatus struct {

	// conditions
	Conditions []*V1alpha1ApplicationCondition `json:"conditions"`

	// health
	Health *V1alpha1HealthStatus `json:"health,omitempty"`

	// history
	History []*V1alpha1RevisionHistory `json:"history"`

	// observed at
	// Format: date-time
	ObservedAt strfmt.DateTime `json:"observedAt,omitempty"`

	// operation state
	OperationState *V1alpha1OperationState `json:"operationState,omitempty"`

	// reconciled at
	// Format: date-time
	ReconciledAt strfmt.DateTime `json:"reconciledAt,omitempty"`

	// resources
	Resources []*V1alpha1ResourceStatus `json:"resources"`

	// source type
	SourceType string `json:"sourceType,omitempty"`

	// summary
	Summary *V1alpha1ApplicationSummary `json:"summary,omitempty"`

	// sync
	Sync *V1alpha1SyncStatus `json:"sync,omitempty"`
}

// Validate validates this v1alpha1 application status
func (m *V1alpha1ApplicationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObservedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReconciledAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSync(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ApplicationStatus) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateHistory(formats strfmt.Registry) error {

	if swag.IsZero(m.History) { // not required
		return nil
	}

	for i := 0; i < len(m.History); i++ {
		if swag.IsZero(m.History[i]) { // not required
			continue
		}

		if m.History[i] != nil {
			if err := m.History[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("history" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateObservedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ObservedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("observedAt", "body", "date-time", m.ObservedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateOperationState(formats strfmt.Registry) error {

	if swag.IsZero(m.OperationState) { // not required
		return nil
	}

	if m.OperationState != nil {
		if err := m.OperationState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operationState")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateReconciledAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ReconciledAt) { // not required
		return nil
	}

	if err := validate.FormatOf("reconciledAt", "body", "date-time", m.ReconciledAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.Summary) { // not required
		return nil
	}

	if m.Summary != nil {
		if err := m.Summary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationStatus) validateSync(formats strfmt.Registry) error {

	if swag.IsZero(m.Sync) { // not required
		return nil
	}

	if m.Sync != nil {
		if err := m.Sync.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ApplicationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ApplicationStatus) UnmarshalBinary(b []byte) error {
	var res V1alpha1ApplicationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
