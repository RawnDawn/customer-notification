package customer

import "errors"

var ErrInvalidCustomerEmail = errors.New("Invalid customer email")

var ErrCannotSendEmail = errors.New("Cannot send promotional email to the customer")