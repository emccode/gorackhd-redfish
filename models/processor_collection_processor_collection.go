package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
)

/*ProcessorCollectionProcessorCollection processor collection processor collection

swagger:model ProcessorCollection_ProcessorCollection
*/
type ProcessorCollectionProcessorCollection struct {

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

	/* Contains the members of this collection.

	Read Only: true
	*/
	Members []*Odata400IDRef `json:"Members,omitempty"`

	/* members at odata count

	Read Only: true
	*/
	MembersAtOdataCount float64 `json:"Members@odata.count,omitempty"`

	/* members at odata navigation link
	 */
	MembersAtOdataNavigationLink *Odata400IDRef `json:"Members@odata.navigationLink,omitempty"`

	/* The name of the resource or array element.

	Read Only: true
	*/
	Name string `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`
}

// Validate validates this processor collection processor collection
func (m *ProcessorCollectionProcessorCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessorCollectionProcessorCollection) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	return nil
}
