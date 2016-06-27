package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*Manager100GraphicalConnectTypesSupported manager 1 0 0 graphical connect types supported

swagger:model Manager.1.0.0_GraphicalConnectTypesSupported
*/
type Manager100GraphicalConnectTypesSupported string

// for schema
var manager100GraphicalConnectTypesSupportedEnum []interface{}

func (m Manager100GraphicalConnectTypesSupported) validateManager100GraphicalConnectTypesSupportedEnum(path, location string, value Manager100GraphicalConnectTypesSupported) error {
	if manager100GraphicalConnectTypesSupportedEnum == nil {
		var res []Manager100GraphicalConnectTypesSupported
		if err := json.Unmarshal([]byte(`["KVMIP","Oem"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			manager100GraphicalConnectTypesSupportedEnum = append(manager100GraphicalConnectTypesSupportedEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, manager100GraphicalConnectTypesSupportedEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this manager 1 0 0 graphical connect types supported
func (m Manager100GraphicalConnectTypesSupported) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateManager100GraphicalConnectTypesSupportedEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}