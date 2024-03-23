package cc

type cardValidator struct {
	cardNumber string
	isValid    bool
}

func NewCardValidator(cardNumber string) *cardValidator {
	validator := cardValidator{cardNumber, true}
	return &validator
}

func (c *cardValidator) ValidateLength() *cardValidator {
	if len(c.cardNumber) < 9 || len(c.cardNumber) > 19 {
		c.isValid = false
	}
	return c
}

func (c *cardValidator) ValidateLuhn() *cardValidator {
	if !isValidLuhn(c.cardNumber) {
		c.isValid = false
	}
	return c
}

func (c *cardValidator) GetIsValid() bool {
	return c.isValid
}
