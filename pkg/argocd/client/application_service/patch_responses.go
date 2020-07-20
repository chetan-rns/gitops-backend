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

// PatchReader is a Reader for the Patch structure.
type PatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchOK creates a PatchOK with default headers values
func NewPatchOK() *PatchOK {
	return &PatchOK{}
}

/*PatchOK handles this case with default header values.

(empty)
*/
type PatchOK struct {
	Payload *models.V1alpha1Application
}

func (o *PatchOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/applications/{name}][%d] patchOK  %+v", 200, o.Payload)
}

func (o *PatchOK) GetPayload() *models.V1alpha1Application {
	return o.Payload
}

func (o *PatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
