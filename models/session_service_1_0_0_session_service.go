package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*SessionService100SessionService This is the schema definition for the Session Service.  It represents the properties for the service itself and has links to the actual list of sessions.

swagger:model SessionService.1.0.0_SessionService
*/
type SessionService100SessionService struct {

	/* at odata context

	Read Only: true
	*/
	AtOdataContext strfmt.URI `json:"@odata.context,omitempty"`

	/* at odata id

	Read Only: true
	*/
	AtOdataID strfmt.URI `json:"@odata.id,omitempty"`

	/* at odata type

	Read Only: true
	*/
	AtOdataType string `json:"@odata.type,omitempty"`

	/* Provides a description of this resource and is used for commonality  in the schema definitions.

	Read Only: true
	*/
	Description string `json:"Description,omitempty"`

	/* Uniquely identifies the resource within the collection of like resources.

	Read Only: true
	*/
	ID string `json:"Id,omitempty"`

	/* The name of the resource or array element.

	Read Only: true
	*/
	Name string `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* This indicates whether this service is enabled.
	 */
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	/* This is the number of seconds of inactivity that a session may have before the session service closes the session due to inactivity.

	Maximum: 86400
	Minimum: 30
	*/
	SessionTimeout float64 `json:"SessionTimeout,omitempty"`

	/* Link to a collection of Sessions

	Read Only: true
	*/
	Sessions *SessionCollectionSessionCollection `json:"Sessions,omitempty"`

	/* status
	 */
	Status *ResourceStatus `json:"Status,omitempty"`
}

// Validate validates this session service 1 0 0 session service
func (m *SessionService100SessionService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSessionTimeout(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SessionService100SessionService) validateSessionTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.SessionTimeout) { // not required
		return nil
	}

	if err := validate.Minimum("SessionTimeout", "body", float64(m.SessionTimeout), 30, false); err != nil {
		return err
	}

	if err := validate.Maximum("SessionTimeout", "body", float64(m.SessionTimeout), 86400, false); err != nil {
		return err
	}

	return nil
}
