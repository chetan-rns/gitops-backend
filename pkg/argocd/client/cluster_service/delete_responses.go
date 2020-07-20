// Code generated by go-swagger; DO NOT EDIT.

package cluster_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// DeleteReader is a Reader for the Delete structure.
type DeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOK creates a DeleteOK with default headers values
func NewDeleteOK() *DeleteOK {
	return &DeleteOK{}
}

/*DeleteOK handles this case with default header values.

(empty)
*/
type DeleteOK struct {
	Payload models.ClusterClusterResponse
}

func (o *DeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/clusters/{server}][%d] deleteOK  %+v", 200, o.Payload)
}

func (o *DeleteOK) GetPayload() models.ClusterClusterResponse {
	return o.Payload
}

func (o *DeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}