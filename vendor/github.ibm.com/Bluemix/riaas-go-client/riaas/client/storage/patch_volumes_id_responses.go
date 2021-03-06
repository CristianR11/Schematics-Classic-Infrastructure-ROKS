// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PatchVolumesIDReader is a Reader for the PatchVolumesID structure.
type PatchVolumesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchVolumesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchVolumesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchVolumesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchVolumesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPatchVolumesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchVolumesIDOK creates a PatchVolumesIDOK with default headers values
func NewPatchVolumesIDOK() *PatchVolumesIDOK {
	return &PatchVolumesIDOK{}
}

/*PatchVolumesIDOK handles this case with default header values.

dummy
*/
type PatchVolumesIDOK struct {
	Payload *models.Volume
}

func (o *PatchVolumesIDOK) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{id}][%d] patchVolumesIdOK  %+v", 200, o.Payload)
}

func (o *PatchVolumesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVolumesIDBadRequest creates a PatchVolumesIDBadRequest with default headers values
func NewPatchVolumesIDBadRequest() *PatchVolumesIDBadRequest {
	return &PatchVolumesIDBadRequest{}
}

/*PatchVolumesIDBadRequest handles this case with default header values.

error
*/
type PatchVolumesIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchVolumesIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{id}][%d] patchVolumesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchVolumesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVolumesIDNotFound creates a PatchVolumesIDNotFound with default headers values
func NewPatchVolumesIDNotFound() *PatchVolumesIDNotFound {
	return &PatchVolumesIDNotFound{}
}

/*PatchVolumesIDNotFound handles this case with default header values.

error
*/
type PatchVolumesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *PatchVolumesIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{id}][%d] patchVolumesIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchVolumesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVolumesIDInternalServerError creates a PatchVolumesIDInternalServerError with default headers values
func NewPatchVolumesIDInternalServerError() *PatchVolumesIDInternalServerError {
	return &PatchVolumesIDInternalServerError{}
}

/*PatchVolumesIDInternalServerError handles this case with default header values.

error
*/
type PatchVolumesIDInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *PatchVolumesIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /volumes/{id}][%d] patchVolumesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchVolumesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchVolumesIDBody VolumePatch
swagger:model PatchVolumesIDBody
*/
type PatchVolumesIDBody struct {

	// The capacity of the volume in gigabytes
	// Maximum: 64000
	// Minimum: 10
	Capacity int64 `json:"capacity,omitempty"`

	// The bandwidth for the volume
	// Enum: [1000 10000 100000]
	Iops int64 `json:"iops,omitempty"`

	// The user-defined name for this volume
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// profile
	Profile *PatchVolumesIDParamsBodyProfile `json:"profile,omitempty"`
}

// Validate validates this patch volumes ID body
func (o *PatchVolumesIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchVolumesIDBody) validateCapacity(formats strfmt.Registry) error {

	if swag.IsZero(o.Capacity) { // not required
		return nil
	}

	if err := validate.MinimumInt("body"+"."+"capacity", "body", int64(o.Capacity), 10, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("body"+"."+"capacity", "body", int64(o.Capacity), 64000, false); err != nil {
		return err
	}

	return nil
}

var patchVolumesIdBodyTypeIopsPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1000,10000,100000]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchVolumesIdBodyTypeIopsPropEnum = append(patchVolumesIdBodyTypeIopsPropEnum, v)
	}
}

// prop value enum
func (o *PatchVolumesIDBody) validateIopsEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, patchVolumesIdBodyTypeIopsPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PatchVolumesIDBody) validateIops(formats strfmt.Registry) error {

	if swag.IsZero(o.Iops) { // not required
		return nil
	}

	// value enum
	if err := o.validateIopsEnum("body"+"."+"iops", "body", o.Iops); err != nil {
		return err
	}

	return nil
}

func (o *PatchVolumesIDBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (o *PatchVolumesIDBody) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(o.Profile) { // not required
		return nil
	}

	if o.Profile != nil {
		if err := o.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "profile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchVolumesIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchVolumesIDBody) UnmarshalBinary(b []byte) error {
	var res PatchVolumesIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchVolumesIDParamsBodyProfile reference
swagger:model PatchVolumesIDParamsBodyProfile
*/
type PatchVolumesIDParamsBodyProfile struct {

	// The CRN for this snapshot
	Crn string `json:"crn,omitempty"`

	// The user-defined name for this resource
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this patch volumes ID params body profile
func (o *PatchVolumesIDParamsBodyProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchVolumesIDParamsBodyProfile) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"profile"+"."+"name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchVolumesIDParamsBodyProfile) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchVolumesIDParamsBodyProfile) UnmarshalBinary(b []byte) error {
	var res PatchVolumesIDParamsBodyProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
