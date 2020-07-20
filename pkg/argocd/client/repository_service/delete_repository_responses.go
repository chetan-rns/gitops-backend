// Code generated by go-swagger; DO NOT EDIT.

package repository_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// DeleteRepositoryReader is a Reader for the DeleteRepository structure.
type DeleteRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRepositoryOK creates a DeleteRepositoryOK with default headers values
func NewDeleteRepositoryOK() *DeleteRepositoryOK {
	return &DeleteRepositoryOK{}
}

/*DeleteRepositoryOK handles this case with default header values.

(empty)
*/
type DeleteRepositoryOK struct {
	Payload models.RepositoryRepoResponse
}

func (o *DeleteRepositoryOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repositories/{repo}][%d] deleteRepositoryOK  %+v", 200, o.Payload)
}

func (o *DeleteRepositoryOK) GetPayload() models.RepositoryRepoResponse {
	return o.Payload
}

func (o *DeleteRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
