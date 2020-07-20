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

// GetAppDetailsReader is a Reader for the GetAppDetails structure.
type GetAppDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAppDetailsOK creates a GetAppDetailsOK with default headers values
func NewGetAppDetailsOK() *GetAppDetailsOK {
	return &GetAppDetailsOK{}
}

/*GetAppDetailsOK handles this case with default header values.

(empty)
*/
type GetAppDetailsOK struct {
	Payload *models.RepositoryRepoAppDetailsResponse
}

func (o *GetAppDetailsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/repositories/{source.repoURL}/appdetails][%d] getAppDetailsOK  %+v", 200, o.Payload)
}

func (o *GetAppDetailsOK) GetPayload() *models.RepositoryRepoAppDetailsResponse {
	return o.Payload
}

func (o *GetAppDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RepositoryRepoAppDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
