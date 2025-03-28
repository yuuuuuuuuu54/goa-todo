// Code generated by goa v3.20.0, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen backend/design

package server

import (
	user "backend/gen/user"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "user" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// ユーザー名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// メールアドレス
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// パスワード
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// LoginRequestBody is the type of the "user" service "login" endpoint HTTP
// request body.
type LoginRequestBody struct {
	// メールアドレス
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// パスワード
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
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
	ID int `form:"id" json:"id" xml:"id"`
	// ユーザー名
	Name string `form:"name" json:"name" xml:"name"`
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// 作成日時
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// 更新日時
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// LoginResponseBody is the type of the "user" service "login" endpoint HTTP
// response body.
type LoginResponseBody struct {
	// userのid
	ID int `form:"id" json:"id" xml:"id"`
	// ユーザー名
	Name string `form:"name" json:"name" xml:"name"`
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// 作成日時
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// 更新日時
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// GetResponseBody is the type of the "user" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// userのid
	ID int `form:"id" json:"id" xml:"id"`
	// ユーザー名
	Name string `form:"name" json:"name" xml:"name"`
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// 作成日時
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// 更新日時
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// UpdateResponseBody is the type of the "user" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	// userのid
	ID int `form:"id" json:"id" xml:"id"`
	// ユーザー名
	Name string `form:"name" json:"name" xml:"name"`
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// 作成日時
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// 更新日時
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// CreateBadRequestResponseBody is the type of the "user" service "create"
// endpoint HTTP response body for the "BadRequest" error.
type CreateBadRequestResponseBody struct {
	// エラー名
	Name string `form:"name" json:"name" xml:"name"`
	// エラーメッセージ
	Message string `form:"message" json:"message" xml:"message"`
}

// LoginUnauthorizedResponseBody is the type of the "user" service "login"
// endpoint HTTP response body for the "Unauthorized" error.
type LoginUnauthorizedResponseBody struct {
	// エラー名
	Name string `form:"name" json:"name" xml:"name"`
	// エラーメッセージ
	Message string `form:"message" json:"message" xml:"message"`
}

// LoginBadRequestResponseBody is the type of the "user" service "login"
// endpoint HTTP response body for the "BadRequest" error.
type LoginBadRequestResponseBody struct {
	// エラー名
	Name string `form:"name" json:"name" xml:"name"`
	// エラーメッセージ
	Message string `form:"message" json:"message" xml:"message"`
}

// GetNotFoundResponseBody is the type of the "user" service "get" endpoint
// HTTP response body for the "NotFound" error.
type GetNotFoundResponseBody struct {
	// エラー名
	Name string `form:"name" json:"name" xml:"name"`
	// エラーメッセージ
	Message string `form:"message" json:"message" xml:"message"`
}

// UpdateNotFoundResponseBody is the type of the "user" service "update"
// endpoint HTTP response body for the "NotFound" error.
type UpdateNotFoundResponseBody struct {
	// エラー名
	Name string `form:"name" json:"name" xml:"name"`
	// エラーメッセージ
	Message string `form:"message" json:"message" xml:"message"`
}

// UpdateBadRequestResponseBody is the type of the "user" service "update"
// endpoint HTTP response body for the "BadRequest" error.
type UpdateBadRequestResponseBody struct {
	// エラー名
	Name string `form:"name" json:"name" xml:"name"`
	// エラーメッセージ
	Message string `form:"message" json:"message" xml:"message"`
}

// DeleteNotFoundResponseBody is the type of the "user" service "delete"
// endpoint HTTP response body for the "NotFound" error.
type DeleteNotFoundResponseBody struct {
	// エラー名
	Name string `form:"name" json:"name" xml:"name"`
	// エラーメッセージ
	Message string `form:"message" json:"message" xml:"message"`
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "user" service.
func NewCreateResponseBody(res *user.UserResult) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
	return body
}

// NewLoginResponseBody builds the HTTP response body from the result of the
// "login" endpoint of the "user" service.
func NewLoginResponseBody(res *user.UserResult) *LoginResponseBody {
	body := &LoginResponseBody{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
	return body
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "user" service.
func NewGetResponseBody(res *user.UserResult) *GetResponseBody {
	body := &GetResponseBody{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "user" service.
func NewUpdateResponseBody(res *user.UserResult) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
	return body
}

// NewCreateBadRequestResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "user" service.
func NewCreateBadRequestResponseBody(res *user.APIErrorResult) *CreateBadRequestResponseBody {
	body := &CreateBadRequestResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewLoginUnauthorizedResponseBody builds the HTTP response body from the
// result of the "login" endpoint of the "user" service.
func NewLoginUnauthorizedResponseBody(res *user.APIErrorResult) *LoginUnauthorizedResponseBody {
	body := &LoginUnauthorizedResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewLoginBadRequestResponseBody builds the HTTP response body from the result
// of the "login" endpoint of the "user" service.
func NewLoginBadRequestResponseBody(res *user.APIErrorResult) *LoginBadRequestResponseBody {
	body := &LoginBadRequestResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewGetNotFoundResponseBody builds the HTTP response body from the result of
// the "get" endpoint of the "user" service.
func NewGetNotFoundResponseBody(res *user.APIErrorResult) *GetNotFoundResponseBody {
	body := &GetNotFoundResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "user" service.
func NewUpdateNotFoundResponseBody(res *user.APIErrorResult) *UpdateNotFoundResponseBody {
	body := &UpdateNotFoundResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewUpdateBadRequestResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "user" service.
func NewUpdateBadRequestResponseBody(res *user.APIErrorResult) *UpdateBadRequestResponseBody {
	body := &UpdateBadRequestResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewDeleteNotFoundResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "user" service.
func NewDeleteNotFoundResponseBody(res *user.APIErrorResult) *DeleteNotFoundResponseBody {
	body := &DeleteNotFoundResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewCreateUser builds a user service create endpoint payload.
func NewCreateUser(body *CreateRequestBody) *user.User {
	v := &user.User{
		Name:     *body.Name,
		Email:    *body.Email,
		Password: *body.Password,
	}

	return v
}

// NewLoginPayload builds a user service login endpoint payload.
func NewLoginPayload(body *LoginRequestBody) *user.LoginPayload {
	v := &user.LoginPayload{
		Email:    *body.Email,
		Password: *body.Password,
	}

	return v
}

// NewGetPayload builds a user service get endpoint payload.
func NewGetPayload(userID int) *user.GetPayload {
	v := &user.GetPayload{}
	v.UserID = userID

	return v
}

// NewUpdatePayload builds a user service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, userID int) *user.UpdatePayload {
	v := &user.UpdatePayload{
		Name:  body.Name,
		Email: body.Email,
	}
	v.UserID = userID

	return v
}

// NewDeletePayload builds a user service delete endpoint payload.
func NewDeletePayload(userID int) *user.DeletePayload {
	v := &user.DeletePayload{}
	v.UserID = userID

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 3, true))
		}
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	if body.Password != nil {
		if utf8.RuneCountInString(*body.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", *body.Password, utf8.RuneCountInString(*body.Password), 8, true))
		}
	}
	return
}

// ValidateLoginRequestBody runs the validations defined on LoginRequestBody
func ValidateLoginRequestBody(body *LoginRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 3, true))
		}
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	return
}
