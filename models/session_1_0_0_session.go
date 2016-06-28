package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Session100Session This is the base type for resources and referenceable members.

swagger:model Session.1.0.0_Session
*/
type Session100Session struct {

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

	/* This property is used in a POST to specify a password when creating a new session.  This property is null on a GET.
	 */
	Password string `json:"Password,omitempty"`

	/* user name

	Read Only: true
	*/
	UserName string `json:"UserName,omitempty"`
}

// Validate validates this session 1 0 0 session
func (m *Session100Session) Validate(formats strfmt.Registry) error {
	return nil
}
