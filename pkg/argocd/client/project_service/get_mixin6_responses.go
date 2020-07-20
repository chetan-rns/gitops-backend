// Code generated by go-swagger; DO NOT EDIT.

package project_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// GetMixin6Reader is a Reader for the GetMixin6 structure.
type GetMixin6Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMixin6Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMixin6OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMixin6OK creates a GetMixin6OK with default headers values
func NewGetMixin6OK() *GetMixin6OK {
	return &GetMixin6OK{}
}

/*GetMixin6OK handles this case with default header values.

(empty)
*/
type GetMixin6OK struct {
	Payload *models.V1alpha1AppProject
}

func (o *GetMixin6OK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{name}][%d] getMixin6OK  %+v", 200, o.Payload)
}

func (o *GetMixin6OK) GetPayload() *models.V1alpha1AppProject {
	return o.Payload
}

func (o *GetMixin6OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1AppProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
