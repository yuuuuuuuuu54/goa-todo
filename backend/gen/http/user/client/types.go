// Code generated by goa v3.20.0, DO NOT EDIT.
//
// user HTTP client types
//
// Command:
// $ goa gen backend/design

package client

import (
	user "backend/gen/user"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "user" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// ユーザー名
	Name string `form:"name" json:"name" xml:"name"`
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// パスワード
	Password string `form:"password" json:"password" xml:"password"`
}

// LoginRequestBody is the type of the "user" service "login" endpoint HTTP
// request body.
type LoginRequestBody struct {
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// パスワード
	Password string `form:"password" json:"password" xml:"password"`
}

// UpdateRequestBody is the type of the "user" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// ユーザー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// メールアドレス
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// CreateResponseBody is the type of the "user" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// userのid
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// ユーザー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// メールアドレス
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// 作成日時
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// 更新日時
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// LoginResponseBody is the type of the "user" service "login" endpoint HTTP
// response body.
type LoginResponseBody struct {
	// userのid
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// ユーザー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// メールアドレス
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// 作成日時
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// 更新日時
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// GetResponseBody is the type of the "user" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// userのid
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// ユーザー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// メールアドレス
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// 作成日時
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// 更新日時
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// UpdateResponseBody is the type of the "user" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	// userのid
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// ユーザー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// メールアドレス
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// 作成日時
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// 更新日時
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// CreateBadRequestResponseBody is the type of the "user" service "create"
// endpoint HTTP response body for the "BadRequest" error.
type CreateBadRequestResponseBody struct {
	// エラー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// エラーメッセージ
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// LoginUnauthorizedResponseBody is the type of the "user" service "login"
// endpoint HTTP response body for the "Unauthorized" error.
type LoginUnauthorizedResponseBody struct {
	// エラー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// エラーメッセージ
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// LoginBadRequestResponseBody is the type of the "user" service "login"
// endpoint HTTP response body for the "BadRequest" error.
type LoginBadRequestResponseBody struct {
	// エラー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// エラーメッセージ
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// GetNotFoundResponseBody is the type of the "user" service "get" endpoint
// HTTP response body for the "NotFound" error.
type GetNotFoundResponseBody struct {
	// エラー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// エラーメッセージ
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateNotFoundResponseBody is the type of the "user" service "update"
// endpoint HTTP response body for the "NotFound" error.
type UpdateNotFoundResponseBody struct {
	// エラー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// エラーメッセージ
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// UpdateBadRequestResponseBody is the type of the "user" service "update"
// endpoint HTTP response body for the "BadRequest" error.
type UpdateBadRequestResponseBody struct {
	// エラー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// エラーメッセージ
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteNotFoundResponseBody is the type of the "user" service "delete"
// endpoint HTTP response body for the "NotFound" error.
type DeleteNotFoundResponseBody struct {
	// エラー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// エラーメッセージ
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "user" service.
func NewCreateRequestBody(p *user.User) *CreateRequestBody {
	body := &CreateRequestBody{
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
	}
	return body
}

// NewLoginRequestBody builds the HTTP request body from the payload of the
// "login" endpoint of the "user" service.
func NewLoginRequestBody(p *user.LoginPayload) *LoginRequestBody {
	body := &LoginRequestBody{
		Email:    p.Email,
		Password: p.Password,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "user" service.
func NewUpdateRequestBody(p *user.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Name:  p.Name,
		Email: p.Email,
	}
	return body
}

// NewCreateUserResultCreated builds a "user" service "create" endpoint result
// from a HTTP "Created" response.
func NewCreateUserResultCreated(body *CreateResponseBody) *user.UserResult {
	v := &user.UserResult{
		ID:        *body.ID,
		Name:      *body.Name,
		Email:     *body.Email,
		CreatedAt: *body.CreatedAt,
		UpdatedAt: *body.UpdatedAt,
	}

	return v
}

// NewCreateBadRequest builds a user service create endpoint BadRequest error.
func NewCreateBadRequest(body *CreateBadRequestResponseBody) *user.APIErrorResult {
	v := &user.APIErrorResult{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// NewLoginUserResultOK builds a "user" service "login" endpoint result from a
// HTTP "OK" response.
func NewLoginUserResultOK(body *LoginResponseBody) *user.UserResult {
	v := &user.UserResult{
		ID:        *body.ID,
		Name:      *body.Name,
		Email:     *body.Email,
		CreatedAt: *body.CreatedAt,
		UpdatedAt: *body.UpdatedAt,
	}

	return v
}

// NewLoginUnauthorized builds a user service login endpoint Unauthorized error.
func NewLoginUnauthorized(body *LoginUnauthorizedResponseBody) *user.APIErrorResult {
	v := &user.APIErrorResult{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// NewLoginBadRequest builds a user service login endpoint BadRequest error.
func NewLoginBadRequest(body *LoginBadRequestResponseBody) *user.APIErrorResult {
	v := &user.APIErrorResult{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// NewGetUserResultOK builds a "user" service "get" endpoint result from a HTTP
// "OK" response.
func NewGetUserResultOK(body *GetResponseBody) *user.UserResult {
	v := &user.UserResult{
		ID:        *body.ID,
		Name:      *body.Name,
		Email:     *body.Email,
		CreatedAt: *body.CreatedAt,
		UpdatedAt: *body.UpdatedAt,
	}

	return v
}

// NewGetNotFound builds a user service get endpoint NotFound error.
func NewGetNotFound(body *GetNotFoundResponseBody) *user.APIErrorResult {
	v := &user.APIErrorResult{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// NewUpdateUserResultOK builds a "user" service "update" endpoint result from
// a HTTP "OK" response.
func NewUpdateUserResultOK(body *UpdateResponseBody) *user.UserResult {
	v := &user.UserResult{
		ID:        *body.ID,
		Name:      *body.Name,
		Email:     *body.Email,
		CreatedAt: *body.CreatedAt,
		UpdatedAt: *body.UpdatedAt,
	}

	return v
}

// NewUpdateNotFound builds a user service update endpoint NotFound error.
func NewUpdateNotFound(body *UpdateNotFoundResponseBody) *user.APIErrorResult {
	v := &user.APIErrorResult{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// NewUpdateBadRequest builds a user service update endpoint BadRequest error.
func NewUpdateBadRequest(body *UpdateBadRequestResponseBody) *user.APIErrorResult {
	v := &user.APIErrorResult{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// NewDeleteNotFound builds a user service delete endpoint NotFound error.
func NewDeleteNotFound(body *DeleteNotFoundResponseBody) *user.APIErrorResult {
	v := &user.APIErrorResult{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// ValidateCreateResponseBody runs the validations defined on CreateResponseBody
func ValidateCreateResponseBody(body *CreateResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "body"))
	}
	if body.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "body"))
	}
	return
}

// ValidateLoginResponseBody runs the validations defined on LoginResponseBody
func ValidateLoginResponseBody(body *LoginResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "body"))
	}
	if body.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "body"))
	}
	return
}

// ValidateGetResponseBody runs the validations defined on GetResponseBody
func ValidateGetResponseBody(body *GetResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "body"))
	}
	if body.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "body"))
	}
	return
}

// ValidateUpdateResponseBody runs the validations defined on UpdateResponseBody
func ValidateUpdateResponseBody(body *UpdateResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "body"))
	}
	if body.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "body"))
	}
	return
}

// ValidateCreateBadRequestResponseBody runs the validations defined on
// create_BadRequest_response_body
func ValidateCreateBadRequestResponseBody(body *CreateBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateLoginUnauthorizedResponseBody runs the validations defined on
// login_Unauthorized_response_body
func ValidateLoginUnauthorizedResponseBody(body *LoginUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateLoginBadRequestResponseBody runs the validations defined on
// login_BadRequest_response_body
func ValidateLoginBadRequestResponseBody(body *LoginBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateGetNotFoundResponseBody runs the validations defined on
// get_NotFound_response_body
func ValidateGetNotFoundResponseBody(body *GetNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateNotFoundResponseBody runs the validations defined on
// update_NotFound_response_body
func ValidateUpdateNotFoundResponseBody(body *UpdateNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateUpdateBadRequestResponseBody runs the validations defined on
// update_BadRequest_response_body
func ValidateUpdateBadRequestResponseBody(body *UpdateBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDeleteNotFoundResponseBody runs the validations defined on
// delete_NotFound_response_body
func ValidateDeleteNotFoundResponseBody(body *DeleteNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}
