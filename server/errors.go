package server

import (
	"fmt"
)

// TransactionFailureType to represent different types of transaction errors.
type TransactionFailureType string

const (
	InsufficientFunds TransactionFailureType = "Insufficient Funds"
	AccountLocked      TransactionFailureType = "Account Locked"
	DailyLimitExceeded TransactionFailureType = "Daily Limit Exceeded"
)

// TransactionError is a custom error type for transaction failures.
type TransactionError struct {
	FailureType TransactionFailureType 
	AccountID   string                 
	Details     string                 
	Err         error                
}

func (te *TransactionError) Error() string {
	return fmt.Sprintf("Transaction Error: %s (Account: %s) - %s", te.FailureType, te.AccountID, te.Details)
}

// Unwrap allows TransactionError to support error wrapping.
func (te *TransactionError) Unwrap() error {
	return te.Err
}

func NewTransactionError(failureType TransactionFailureType, accountID, details string, err error) error {
	return &TransactionError{
		FailureType: failureType,
		AccountID:   accountID,
		Details:     details,
		Err:         err,
	}
}

func HandleError(err error) {
	switch e := err.(type) {
	case *TransactionError:
		fmt.Printf("Handling Transaction Error:\n- Type: %s\n- Account: %s\n- Details: %s\n\n", e.FailureType, e.AccountID, e.Details)
		if e.Err != nil {
			fmt.Printf("Wrapped Error: %v\n", e.Err)
		}
	default:
		fmt.Println("General Error:", err)
	}
}