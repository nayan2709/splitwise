package expense

// Operation constants
const (
	MoveIn   = "MOVE_IN"
	Spend    = "SPEND"
	Dues     = "DUES"
	ClearDue = "CLEAR_DUE"
	MoveOut  = "MOVE_OUT"
)

// Response constants
const (
	Success          = "SUCCESS"
	Houseful         = "HOUSEFUL"
	MemberNotFound   = "MEMBER_NOT_FOUND"
	IncorrectPayment = "INCORRECT_PAYMENT"
	Failure          = "FAILURE"

	// Error messages
	InvalidOperation = "INVALID_OPERATION"
	InvalidAmount    = "INVALID_AMOUNT"
)
