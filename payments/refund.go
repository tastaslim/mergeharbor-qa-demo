// Package payments handles charges and refunds for the QA demo service.
package payments

import "errors"

// ErrRefundTooLarge is returned when a refund exceeds what is left of the charge.
var ErrRefundTooLarge = errors.New("refund exceeds the refundable amount")

// Charge is a captured card payment, in cents.
type Charge struct {
	ID       string
	Amount   int64
	Refunded int64
}

// Gateway sends refunds to the payment processor.
type Gateway interface {
	Refund(chargeID string, amount int64) error
}

// Refund returns amount cents of the charge through the gateway.
func Refund(gw Gateway, c *Charge, amount int64) error {
	if amount <= 0 {
		return errors.New("refund amount must be positive")
	}
	if amount > c.Amount-c.Refunded {
		return ErrRefundTooLarge
	}
	if err := gw.Refund(c.ID, amount); err != nil {
		return err
	}
	c.Refunded += amount
	return nil
}
