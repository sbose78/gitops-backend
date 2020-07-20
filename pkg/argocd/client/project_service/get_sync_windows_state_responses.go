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

// GetSyncWindowsStateReader is a Reader for the GetSyncWindowsState structure.
type GetSyncWindowsStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSyncWindowsStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSyncWindowsStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSyncWindowsStateOK creates a GetSyncWindowsStateOK with default headers values
func NewGetSyncWindowsStateOK() *GetSyncWindowsStateOK {
	return &GetSyncWindowsStateOK{}
}

/*GetSyncWindowsStateOK handles this case with default header values.

(empty)
*/
type GetSyncWindowsStateOK struct {
	Payload *models.ProjectSyncWindowsResponse
}

func (o *GetSyncWindowsStateOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{name}/syncwindows][%d] getSyncWindowsStateOK  %+v", 200, o.Payload)
}

func (o *GetSyncWindowsStateOK) GetPayload() *models.ProjectSyncWindowsResponse {
	return o.Payload
}

func (o *GetSyncWindowsStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectSyncWindowsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}