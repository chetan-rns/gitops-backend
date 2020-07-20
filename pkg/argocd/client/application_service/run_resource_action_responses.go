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

// RunResourceActionReader is a Reader for the RunResourceAction structure.
type RunResourceActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunResourceActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunResourceActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRunResourceActionOK creates a RunResourceActionOK with default headers values
func NewRunResourceActionOK() *RunResourceActionOK {
	return &RunResourceActionOK{}
}

/*RunResourceActionOK handles this case with default header values.

(empty)
*/
type RunResourceActionOK struct {
	Payload models.ApplicationApplicationResponse
}

func (o *RunResourceActionOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/applications/{name}/resource/actions][%d] runResourceActionOK  %+v", 200, o.Payload)
}

func (o *RunResourceActionOK) GetPayload() models.ApplicationApplicationResponse {
	return o.Payload
}

func (o *RunResourceActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
