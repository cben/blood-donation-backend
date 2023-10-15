// Package bloodinfo provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package bloodinfo

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for UserRole.
const (
	Admin    UserRole = "Admin"
	Reporter UserRole = "Reporter"
)

// ApiError defines model for ApiError.
type ApiError struct {
	Message string `json:"message"`
}

// Station defines model for Station.
type Station struct {
	Address   string             `json:"address"`
	CloseTime openapi_types.Date `json:"close_time"`

	// DeletedAt swagger:ignore
	DeletedAt interface{}        `json:"deleted_at"`
	Id        int64              `gorm:"primaryKey" json:"id"`
	IsOpen    bool               `json:"is_open"`
	OpenTime  openapi_types.Date `json:"open_time"`
}

// User defines model for User.
type User struct {
	// DeletedAt swagger:ignore
	DeletedAt   interface{} `json:"deleted_at"`
	Description string      `json:"description"`
	Email       string      `json:"email"`
	FirstName   string      `json:"first_name"`
	Id          int64       `gorm:"primaryKey" json:"id"`
	LastName    string      `json:"last_name"`
	Phone       string      `json:"phone"`
	Role        UserRole    `json:"role"`
}

// UserRole defines model for User.Role.
type UserRole string

// UpdateStationJSONBody defines parameters for UpdateStation.
type UpdateStationJSONBody struct {
	// IsOpen New status for station's open status
	IsOpen *bool `json:"isOpen,omitempty"`
}

// UpdateStationJSONRequestBody defines body for UpdateStation for application/json ContentType.
type UpdateStationJSONRequestBody UpdateStationJSONBody

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = User

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all stations
	// (GET /stations)
	GetStations(w http.ResponseWriter, r *http.Request)
	// Update station
	// (PUT /stations/{stationId})
	UpdateStation(w http.ResponseWriter, r *http.Request, stationId int64)
	// Get all users
	// (GET /users)
	GetUsers(w http.ResponseWriter, r *http.Request)
	// Create user.go
	// (POST /users)
	CreateUser(w http.ResponseWriter, r *http.Request)
	// Delete user.go
	// (DELETE /users/{id})
	DeleteUser(w http.ResponseWriter, r *http.Request, id int)
	// Update user.go
	// (PUT /users/{id})
	UpdateUser(w http.ResponseWriter, r *http.Request, id int64)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Get all stations
// (GET /stations)
func (_ Unimplemented) GetStations(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update station
// (PUT /stations/{stationId})
func (_ Unimplemented) UpdateStation(w http.ResponseWriter, r *http.Request, stationId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get all users
// (GET /users)
func (_ Unimplemented) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create user.go
// (POST /users)
func (_ Unimplemented) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete user.go
// (DELETE /users/{id})
func (_ Unimplemented) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update user.go
// (PUT /users/{id})
func (_ Unimplemented) UpdateUser(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetStations operation middleware
func (siw *ServerInterfaceWrapper) GetStations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetStations(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateStation operation middleware
func (siw *ServerInterfaceWrapper) UpdateStation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "stationId" -------------
	var stationId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "stationId", runtime.ParamLocationPath, chi.URLParam(r, "stationId"), &stationId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "stationId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateStation(w, r, stationId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUsers operation middleware
func (siw *ServerInterfaceWrapper) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsers(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUser(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUser(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/stations", wrapper.GetStations)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/stations/{stationId}", wrapper.UpdateStation)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users", wrapper.GetUsers)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/users", wrapper.CreateUser)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/users/{id}", wrapper.DeleteUser)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/users/{id}", wrapper.UpdateUser)
	})

	return r
}

type GetStationsRequestObject struct {
}

type GetStationsResponseObject interface {
	VisitGetStationsResponse(w http.ResponseWriter) error
}

type GetStations200JSONResponse []Station

func (response GetStations200JSONResponse) VisitGetStationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateStationRequestObject struct {
	StationId int64 `json:"stationId"`
	Body      *UpdateStationJSONRequestBody
}

type UpdateStationResponseObject interface {
	VisitUpdateStationResponse(w http.ResponseWriter) error
}

type UpdateStation200JSONResponse Station

func (response UpdateStation200JSONResponse) VisitUpdateStationResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateStation400JSONResponse ApiError

func (response UpdateStation400JSONResponse) VisitUpdateStationResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type UpdateStation401JSONResponse ApiError

func (response UpdateStation401JSONResponse) VisitUpdateStationResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type UpdateStation403JSONResponse ApiError

func (response UpdateStation403JSONResponse) VisitUpdateStationResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type UpdateStation404JSONResponse ApiError

func (response UpdateStation404JSONResponse) VisitUpdateStationResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type UpdateStation405JSONResponse ApiError

func (response UpdateStation405JSONResponse) VisitUpdateStationResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(405)

	return json.NewEncoder(w).Encode(response)
}

type UpdateStation409JSONResponse ApiError

func (response UpdateStation409JSONResponse) VisitUpdateStationResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type UpdateStation500JSONResponse ApiError

func (response UpdateStation500JSONResponse) VisitUpdateStationResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse []User

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUsers401JSONResponse ApiError

func (response GetUsers401JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetUsers500JSONResponse ApiError

func (response GetUsers500JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type CreateUserRequestObject struct {
	Body *CreateUserJSONRequestBody
}

type CreateUserResponseObject interface {
	VisitCreateUserResponse(w http.ResponseWriter) error
}

type CreateUser201JSONResponse User

func (response CreateUser201JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CreateUser401JSONResponse ApiError

func (response CreateUser401JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type CreateUser500JSONResponse ApiError

func (response CreateUser500JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteUserRequestObject struct {
	Id int `json:"id"`
}

type DeleteUserResponseObject interface {
	VisitDeleteUserResponse(w http.ResponseWriter) error
}

type DeleteUser200Response struct {
}

func (response DeleteUser200Response) VisitDeleteUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteUser401JSONResponse ApiError

func (response DeleteUser401JSONResponse) VisitDeleteUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type DeleteUser500JSONResponse ApiError

func (response DeleteUser500JSONResponse) VisitDeleteUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type UpdateUserRequestObject struct {
	Id   int64 `json:"id"`
	Body *UpdateUserJSONRequestBody
}

type UpdateUserResponseObject interface {
	VisitUpdateUserResponse(w http.ResponseWriter) error
}

type UpdateUser200JSONResponse User

func (response UpdateUser200JSONResponse) VisitUpdateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateUser401JSONResponse ApiError

func (response UpdateUser401JSONResponse) VisitUpdateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type UpdateUser500JSONResponse ApiError

func (response UpdateUser500JSONResponse) VisitUpdateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all stations
	// (GET /stations)
	GetStations(ctx context.Context, request GetStationsRequestObject) (GetStationsResponseObject, error)
	// Update station
	// (PUT /stations/{stationId})
	UpdateStation(ctx context.Context, request UpdateStationRequestObject) (UpdateStationResponseObject, error)
	// Get all users
	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
	// Create user.go
	// (POST /users)
	CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error)
	// Delete user.go
	// (DELETE /users/{id})
	DeleteUser(ctx context.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error)
	// Update user.go
	// (PUT /users/{id})
	UpdateUser(ctx context.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHttpHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHttpMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetStations operation middleware
func (sh *strictHandler) GetStations(w http.ResponseWriter, r *http.Request) {
	var request GetStationsRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetStations(ctx, request.(GetStationsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetStations")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetStationsResponseObject); ok {
		if err := validResponse.VisitGetStationsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdateStation operation middleware
func (sh *strictHandler) UpdateStation(w http.ResponseWriter, r *http.Request, stationId int64) {
	var request UpdateStationRequestObject

	request.StationId = stationId

	var body UpdateStationJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateStation(ctx, request.(UpdateStationRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateStation")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(UpdateStationResponseObject); ok {
		if err := validResponse.VisitUpdateStationResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var request GetUsersRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx, request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		if err := validResponse.VisitGetUsersResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateUser operation middleware
func (sh *strictHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request CreateUserRequestObject

	var body CreateUserJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.CreateUser(ctx, request.(CreateUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateUser")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(CreateUserResponseObject); ok {
		if err := validResponse.VisitCreateUserResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteUser operation middleware
func (sh *strictHandler) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	var request DeleteUserRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUser(ctx, request.(DeleteUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUser")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteUserResponseObject); ok {
		if err := validResponse.VisitDeleteUserResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdateUser operation middleware
func (sh *strictHandler) UpdateUser(w http.ResponseWriter, r *http.Request, id int64) {
	var request UpdateUserRequestObject

	request.Id = id

	var body UpdateUserJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateUser(ctx, request.(UpdateUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateUser")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(UpdateUserResponseObject); ok {
		if err := validResponse.VisitUpdateUserResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RX32/bNhD+V4TbgL0otrOkQ6c9pcs6BN2aYkGe2iCgzbPMguJxJJUfDfS/DyQlxbKY",
	"H9tq9yFPlnXH+473fTye7mBBlSaFylko7sAuVlix8HikxW/GkPHP2pBG4wQGS4XWshL9o7vVCAVYZ4Qq",
	"oWlyMPh3LQxyKD72jhd550jzz7hw0ORw5pgTpMbRGecGrU1Ez2EhyeKlE1UAX5KpmIMCOHMI+dido0SH",
	"/NI73QFHuzBCR1Sw16ws0RSiVGT86pu9kvYWtXVU7TlWQgElmar4FMIWHtM6VunsWrhV5v9mX0jhL0Jx",
	"vPkEIYCqpWRziVAsmbTY5CD4IFGh3E+H95kK5bBEE9YS02JvQRxLVHt44wzzWYQy+DygAG1ExcztO7wN",
	"hRb2kjSqtTrNiSQy5XfuLc+t0wZpgkPek7AeaVD+e/xBmVNMn1tMiOjbczPATIgNKyZk0rIUxrpLxSpM",
	"mrdIumSPAesVqbTFkAwGVHXlST7ilfDU/YWajEOzRtxjsghhBvtfT6krWZfIsMRP6CRIWi1prIajDyfZ",
	"kkxWMcVKocpsLol4xkmFDpLZ2ElsxhTPaovGTny1hfNbhjfB+bhzPvpwAjlcobEx+P5kNpl1J4ZpAQUc",
	"TGaTA78J5laBimmHEHjBhGB/R5cxKftU4rkx4c8Jjw5n9zaDVpOy8SD8OJv5nwUphyrEZlpLsQje0882",
	"qjM2Zv8kHFZh4fcGl1DAd9P7Fj5t+/e0665NX2dmDLuNZR7mfvoucG3rygstvZkoyo/QvoILv6Ivy/Su",
	"fTrhTTjodaJE59p3ny7oqEDRfNZbNTOsQofG427Gat2yk2Pfh/wbzxXkEE8G9OnAuoadqTFfK+RTB7Rp",
	"LuJytO4N8dt/xdKw2Ql72rbq4Ube43WoSG2Dwtu8f7CZV2NruU+s7+8bR3MNK32uhjVo/qcAn6W7tM5y",
	"OPyKYP2EkkALhklE3N8J4rlitVuREV+Qt8AHOwF+S2YuOEfVoh7uBPU9uewt1arb66udoP6JbkU88+BH",
	"UtJ1X+qfd6yqVzvW8aBDj3ppuj+Hu/DJOyt6JS6s89aw/dsqTIjPvKq+zYHeOeUP3MkdWR3h/v+kJLjw",
	"0x/ZBM2/GvRa6Rw3eY7mQMB/v+ye5vY5t9D+FjBTteAvVUQjKYxV1LeN6Z2Iw1wc2se6Og7vH9RVNLe6",
	"enSY8z4PTnLi8REuPbGN21UCsf0YmbxUMYz4S7aURyb5h5iP5m0zv9Xh/ev0s9nW+9nLvQ9HGky0Mr8A",
	"zVWnvdpIKGDlnLbQ5Jta/IMWTGYcr1CSrlC5LK6FfG1hMZ1K77ci64rXs9cz8CproTcjnnanwmZsTrVb",
	"/54efqcm0hkt7i7+dmW3zeai+ScAAP//Vyl10cMVAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
