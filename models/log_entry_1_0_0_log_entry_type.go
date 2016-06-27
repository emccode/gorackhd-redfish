package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*LogEntry100LogEntryType log entry 1 0 0 log entry type

swagger:model LogEntry.1.0.0_LogEntryType
*/
type LogEntry100LogEntryType string

// for schema
var logEntry100LogEntryTypeEnum []interface{}

func (m LogEntry100LogEntryType) validateLogEntry100LogEntryTypeEnum(path, location string, value LogEntry100LogEntryType) error {
	if logEntry100LogEntryTypeEnum == nil {
		var res []LogEntry100LogEntryType
		if err := json.Unmarshal([]byte(`["Event","SEL","Oem"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			logEntry100LogEntryTypeEnum = append(logEntry100LogEntryTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, logEntry100LogEntryTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this log entry 1 0 0 log entry type
func (m LogEntry100LogEntryType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLogEntry100LogEntryTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}