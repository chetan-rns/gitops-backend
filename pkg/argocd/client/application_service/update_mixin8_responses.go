// Code generated by go-swagger; DO NOT EDIT.

package application_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// UpdateMixin8Reader is a Reader for the UpdateMixin8 structure.
type UpdateMixin8Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMixin8Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMixin8OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateMixin8OK creates a UpdateMixin8OK with default headers values
func NewUpdateMixin8OK() *UpdateMixin8OK {
	return &UpdateMixin8OK{}
}

/*UpdateMixin8OK handles this case with default header values.

(empty)
*/
type UpdateMixin8OK struct {
	Payload *models.V1alpha1Application
}

func (o *UpdateMixin8OK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/applications/{application.metadata.name}][%d] updateMixin8OK  %+v", 200, o.Payload)
}

func (o *UpdateMixin8OK) GetPayload() *models.V1alpha1Application {
	return o.Payload
}

func (o *UpdateMixin8OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
