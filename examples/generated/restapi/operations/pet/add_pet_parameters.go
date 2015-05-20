package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/casualjim/go-swagger/errors"
	"github.com/casualjim/go-swagger/middleware"

	"github.com/casualjim/go-swagger/examples/generated/models"
)

// AddPetParams contains all the bound params for the add pet operation
// typically these are obtained from a http.Request
type AddPetParams struct {
	// Pet object that needs to be added to the store
	Body models.Pet
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddPetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	if err := route.Consumer.Consume(r.Body, &o.Body); err != nil {
		res = append(res, errors.NewParseError("body", "body", "", err))
	} else {
		if err := o.Body.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
