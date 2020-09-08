// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package openapi

import (
	"net/http"

	"github.com/go-openapi/swag"
	"github.com/powerman/go-service-goswagger-clean-example/api/openapi/model"
	"github.com/powerman/go-service-goswagger-clean-example/api/openapi/restapi/op"
	"github.com/powerman/go-service-goswagger-clean-example/pkg/def"
)

func errAddContact(log Log, err error, code errCode) op.AddContactResponder {
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError { // Do no expose details about internal errors.
		msg = "internal error"
	}

	return op.NewAddContactDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(code.extra),
		Message: swag.String(msg),
	})
}
