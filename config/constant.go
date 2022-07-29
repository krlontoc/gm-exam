package config

const (
	DBError          = "Database Error"
	InternalError    = "Internal Error"
	InvalidTokenType = "Authorization header format must be Bearer {token}."
	Success          = "Success"

	InvalidToken   = "Invalid Token"
	TokenRequired  = "Token Required"
	InvalidPayload = "Invalid payload"
	RequiredField  = "Fill required fields"

	InvalidLogin     = "Invalid Username/Password"
	InvalidRecipient = "Invalid Recipient"

	DepositTransaction  = "Deposit Transaction"
	InsufficientBalance = "Insufficient Balance"
)

func Health() string {
	return ""
}
