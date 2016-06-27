package models

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*ManagerAccountManagerAccount manager account manager account

swagger:model ManagerAccount_ManagerAccount
*/
type ManagerAccountManagerAccount struct {
	Odata400IDRef
}

// Validate validates this manager account manager account
func (m *ManagerAccountManagerAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Odata400IDRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}