package helpers

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) Validate(s any) error {
	v.validate.RegisterValidation("validateUrl", ValidateUrl)
	return v.validate.Struct(s)
}

func ValidateUrl(fl validator.FieldLevel) bool {
	regex := `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
	return regexp.MustCompile(regex).MatchString(fl.Field().String())
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "email":
		return "Invalid email format"
	case "url":
		return "Invalid url format"
	case "min":
		return "Should be at least " + fe.Param() + " characters"
	case "max":
		return "Should be at most " + fe.Param() + " characters"
	case "numeric":
		return "Should be a number"
	case "validateUrl":
		return "Invalid url format"
	}
	return "Unknown error"
}

func ReturnErrorMsg(ctx *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
	}
}

func CustomErrorMsg(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, gin.H{"error": message})
}
