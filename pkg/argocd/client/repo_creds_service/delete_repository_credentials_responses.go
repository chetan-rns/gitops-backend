// Code generated by go-swagger; DO NOT EDIT.

package repo_creds_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// DeleteRepositoryCredentialsReader is a Reader for the DeleteRepositoryCredentials structure.
type DeleteRepositoryCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepositoryCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRepositoryCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRepositoryCredentialsOK creates a DeleteRepositoryCredentialsOK with default headers values
func NewDeleteRepositoryCredentialsOK() *DeleteRepositoryCredentialsOK {
	return &DeleteRepositoryCredentialsOK{}
}

/*DeleteRepositoryCredentialsOK handles this case with default header values.

(empty)
*/
type DeleteRepositoryCredentialsOK struct {
	Payload models.RepocredsRepoCredsResponse
}

func (o *DeleteRepositoryCredentialsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repocreds/{url}][%d] deleteRepositoryCredentialsOK  %+v", 200, o.Payload)
}

func (o *DeleteRepositoryCredentialsOK) GetPayload() models.RepocredsRepoCredsResponse {
	return o.Payload
}

func (o *DeleteRepositoryCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
