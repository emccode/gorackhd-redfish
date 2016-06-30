package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewRemoveAccountParams creates a new RemoveAccountParams object
// with the default values initialized.
func NewRemoveAccountParams() *RemoveAccountParams {
	var ()
	return &RemoveAccountParams{}
}

/*RemoveAccountParams contains all the parameters to send to the API endpoint
for the remove account operation typically these are written to a http.Request
*/
type RemoveAccountParams struct {

	/*Name*/
	Name string
}

// WithName adds the name to the remove account params
func (o *RemoveAccountParams) WithName(name string) *RemoveAccountParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveAccountParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}