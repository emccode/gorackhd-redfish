package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
)

/*Message100Message message 1 0 0 message

swagger:model Message.1.0.0_Message
*/
type Message100Message struct {

	/* This is the human readable message, if provided.
	 */
	Message string `json:"Message,omitempty"`

	/* This array of message arguments are substituted for the arguments in the message when looked up in the message registry.
	 */
	MessageArgs []string `json:"MessageArgs,omitempty"`

	/* This is the key for this message which can be used to look up the message in a message registry.
	 */
	MessageID string `json:"MessageId,omitempty"`

	/* Oem extension object.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* This is an array of properties described by the message.
	 */
	RelatedProperties []string `json:"RelatedProperties,omitempty"`

	/* Used to provide suggestions on how to resolve the situation that caused the error.
	 */
	Resolution string `json:"Resolution,omitempty"`

	/* This is the severity of the errors.
	 */
	Severity string `json:"Severity,omitempty"`
}

// Validate validates this message 1 0 0 message
func (m *Message100Message) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessageArgs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRelatedProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Message100Message) validateMessageArgs(formats strfmt.Registry) error {

	if swag.IsZero(m.MessageArgs) { // not required
		return nil
	}

	return nil
}

func (m *Message100Message) validateRelatedProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedProperties) { // not required
		return nil
	}

	return nil
}