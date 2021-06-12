package viewmodels

// Environment (Live or Test)
type Environment = string

const (
	// NotSelected -
	NotSelected Environment = "none"

	// Live environment
	Live Environment = "live"

	// Test environment
	Test Environment = "test"
)

// Data contains basic user data after user is authorised
type Data struct {
	UserID           int64
	CompanyID        int64
	Email            string
	Active           bool
	Role             string
	UserRoleID       int64
	Environment      Environment
	RequestID        string
	AdditinalLogData []interface{}
}
