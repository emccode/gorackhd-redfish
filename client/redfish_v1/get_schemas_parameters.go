package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetSchemasParams creates a new GetSchemasParams object
// with the default values initialized.
func NewGetSchemasParams() *GetSchemasParams {

	return &GetSchemasParams{}
}

/*GetSchemasParams contains all the parameters to send to the API endpoint
for the get schemas operation typically these are written to a http.Request
*/
type GetSchemasParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchemasParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}