package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/casualjim/go-swagger/examples/generated/models"
	"github.com/casualjim/go-swagger/middleware"
)

// AddPetHandlerFunc turns a function with the right signature into a add pet handler
type AddPetHandlerFunc func(AddPetParams, *models.User) error

func (fn AddPetHandlerFunc) Handle(params AddPetParams, principal *models.User) error {
	return fn(params, principal)
}

// AddPetHandler interface for that can handle valid add pet params
type AddPetHandler interface {
	Handle(AddPetParams, *models.User) error
}

// NewAddPet creates a new http.Handler for the add pet operation
func NewAddPet(ctx *middleware.Context, handler AddPetHandler) *AddPet {
	return &AddPet{Context: ctx, Handler: handler}
}

// AddPet
type AddPet struct {
	Context *middleware.Context
	Params  AddPetParams
	Handler AddPetHandler
}

func (o *AddPet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // it's ok this is really a models.User
	}

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	err = o.Handler.Handle(o.Params, principal) // actually handle the request
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	o.Context.Respond(rw, r, route.Produces, route, nil)

}
