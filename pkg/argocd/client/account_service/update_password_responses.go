// Code generated by go-swagger; DO NOT EDIT.

package account_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/models"
)

// UpdatePasswordReader is a Reader for the UpdatePassword structure.
type UpdatePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePasswordOK creates a UpdatePasswordOK with default headers values
func NewUpdatePasswordOK() *UpdatePasswordOK {
	return &UpdatePasswordOK{}
}

/*UpdatePasswordOK handles this case with default header values.

(empty)
*/
type UpdatePasswordOK struct {
	Payload models.AccountUpdatePasswordResponse
}

func (o *UpdatePasswordOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/account/password][%d] updatePasswordOK  %+v", 200, o.Payload)
}

func (o *UpdatePasswordOK) GetPayload() models.AccountUpdatePasswordResponse {
	return o.Payload
}

func (o *UpdatePasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
