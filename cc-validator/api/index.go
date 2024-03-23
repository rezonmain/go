package handler

import (
	"fmt"
	"net/http"

	"github.com/rezonmain/cc-validator/cc"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cardNumber := r.URL.Query().Get("cn")
	if len(cardNumber) == 0 {
		http.Error(w, "Missing card number, try adding `?cn=xxx` to your request", http.StatusBadRequest)
		return
	}
	isValid := cc.NewCardValidator(cardNumber).ValidateLength().ValidateLuhn().GetIsValid()
	if isValid == false {
		fmt.Fprintf(w, "<h1> Provided card number is <strong>invalid</strong></h1>")
		return
	}
	fmt.Fprintf(w, "<h1>Provided card number is <strong>valid!</strong")
}
