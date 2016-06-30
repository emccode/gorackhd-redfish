package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd-redfish/models"
)

// GetAccountsReader is a Reader for the GetAccounts structure.
type GetAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetAccountsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountsOK creates a GetAccountsOK with default headers values
func NewGetAccountsOK() *GetAccountsOK {
	return &GetAccountsOK{}
}

/*GetAccountsOK handles this case with default header values.

Success
*/
type GetAccountsOK struct {
	Payload *models.ManagerAccountCollectionManagerAccountCollection
}

func (o *GetAccountsOK) Error() string {
	return fmt.Sprintf("[GET /AccountService/Accounts][%d] getAccountsOK  %+v", 200, o.Payload)
}

func (o *GetAccountsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagerAccountCollectionManagerAccountCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsUnauthorized creates a GetAccountsUnauthorized with default headers values
func NewGetAccountsUnauthorized() *GetAccountsUnauthorized {
	return &GetAccountsUnauthorized{}
}

/*GetAccountsUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetAccountsUnauthorized struct {
}

func (o *GetAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /AccountService/Accounts][%d] getAccountsUnauthorized ", 401)
}

func (o *GetAccountsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccountsForbidden creates a GetAccountsForbidden with default headers values
func NewGetAccountsForbidden() *GetAccountsForbidden {
	return &GetAccountsForbidden{}
}

/*GetAccountsForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type GetAccountsForbidden struct {
}

func (o *GetAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /AccountService/Accounts][%d] getAccountsForbidden ", 403)
}

func (o *GetAccountsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccountsInternalServerError creates a GetAccountsInternalServerError with default headers values
func NewGetAccountsInternalServerError() *GetAccountsInternalServerError {
	return &GetAccountsInternalServerError{}
}

/*GetAccountsInternalServerError handles this case with default header values.

Error
*/
type GetAccountsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /AccountService/Accounts][%d] getAccountsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccountsInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}