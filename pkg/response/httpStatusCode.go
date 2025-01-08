package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken = 30001 // token is invalid
	ErrInvalidOTP   = 30002
	ErrSendEmailOtp = 30003
	// User Authentication
	ErrCodeAuthFailed = 40005
	// Register Code
	ErrCodeUserHasExists = 50001 // user has already registered

	// Err Login
	ErrCodeLoginFailed      = 60000
	ErrPasswordIncorrect    = 60006
	ErrCodeOtpNotExists     = 60009
	ErrCodeUserOtpNotExists = 60008
	ErrCodeUserNotExists    = 60007

	// Two Factor Authentication
	ErrCodeTwoFactorAuthSetupFailed  = 80001
	ErrCodeTwoFactorAuthVerifyFailed = 80002
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "token is invalid",
	ErrInvalidOTP:       "Otp error",
	ErrSendEmailOtp:     "Failed to send email OTP",

	ErrCodeUserHasExists: "user has already registered",

	ErrCodeOtpNotExists:     "OTP exists but not registered",
	ErrCodeUserOtpNotExists: "User OTP not exists",
	ErrCodeAuthFailed:       "Authentication failed",
	ErrCodeUserNotExists:    "User not exists",
	ErrCodeLoginFailed:      "Login failed",
	ErrPasswordIncorrect:    "Password is incorrect",

	// Two Factor Authentication
	ErrCodeTwoFactorAuthSetupFailed:  "Two Factor Authentication setup failed",
	ErrCodeTwoFactorAuthVerifyFailed: "Two Factor Authentication verify failed",
}
