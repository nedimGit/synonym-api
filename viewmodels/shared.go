package viewmodels

import (
	"regexp"
	"strconv"
)

// BaseRequest structure that will have all core properties needed for each request
type BaseRequest struct {
	CompanyID int64 // This property is set up on server side and used for additional security check
	Timestamp int64 // This property is set up on server side and used for the time request tracking
}

// BaseResponse structure that will have on all responses from all APIs
type BaseResponse struct {
	Code      int64       `json:"Code"`
	Errors    []Error     `json:"Errors"`
	Data      interface{} `json:"Data"`
	RequestID string      `json:"RequestID"`
}

// Error object that will contain details of the error
type Error struct {
	Code    int64  `json:"error-code"`
	Message string `json:"error-message"`
}

// Token contains all important data for the token
type Token struct {
	Token          string `json:"Token"`
	RefreshToken   string `json:"RefreshToken"`
	ExpirationTime int64  `json:"ExpirationTime"`
}

// Saved is used across the platform
type Saved struct {
	Success bool `json:"Success"`
}

var regx = struct {
	Password *regexp.Regexp
	Email    *regexp.Regexp
}{
	Password: regexp.MustCompile(`^\S{8,20}$`),
	Email:    regexp.MustCompile(`^[a-zA-Z0-9._%\-+]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`),
}

func getBool(s string) (bool, bool) {
	i, err := strconv.ParseBool(s)
	if err != nil {
		return false, false
	}
	return i, true
}

// ValidateEmail will check if email match to regex
func validateEmail(email string) bool {
	return regx.Email.MatchString(email)
}

// ValidatePassword will check if password match to regexp
func validatePassword(password string) bool {
	return regx.Password.MatchString(password)
}
