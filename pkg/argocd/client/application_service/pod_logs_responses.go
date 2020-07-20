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

// PodLogsReader is a Reader for the PodLogs structure.
type PodLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PodLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPodLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPodLogsOK creates a PodLogsOK with default headers values
func NewPodLogsOK() *PodLogsOK {
	return &PodLogsOK{}
}

/*PodLogsOK handles this case with default header values.

(streaming responses)
*/
type PodLogsOK struct {
	Payload *models.ApplicationLogEntry
}

func (o *PodLogsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/applications/{name}/pods/{podName}/logs][%d] podLogsOK  %+v", 200, o.Payload)
}

func (o *PodLogsOK) GetPayload() *models.ApplicationLogEntry {
	return o.Payload
}

func (o *PodLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationLogEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
