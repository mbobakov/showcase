// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/mbobakov/showcase/models"
)

// NewPostDatapointParams creates a new PostDatapointParams object
// no default values defined in spec.
func NewPostDatapointParams() PostDatapointParams {

	return PostDatapointParams{}
}

// PostDatapointParams contains all the bound params for the post datapoint operation
// typically these are obtained from a http.Request
//
// swagger:parameters postDatapoint
type PostDatapointParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Dattime range
	  Required: true
	  In: body
	*/
	Body *models.Datapoint
	/*
	  Required: true
	  Min Length: 1
	  In: path
	*/
	MetricName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostDatapointParams() beforehand.
func (o *PostDatapointParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Datapoint
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	rMetricName, rhkMetricName, _ := route.Params.GetOK("metricName")
	if err := o.bindMetricName(rMetricName, rhkMetricName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMetricName binds and validates parameter MetricName from path.
func (o *PostDatapointParams) bindMetricName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.MetricName = raw

	if err := o.validateMetricName(formats); err != nil {
		return err
	}

	return nil
}

// validateMetricName carries on validations for parameter MetricName
func (o *PostDatapointParams) validateMetricName(formats strfmt.Registry) error {

	if err := validate.MinLength("metricName", "path", o.MetricName, 1); err != nil {
		return err
	}

	return nil
}
