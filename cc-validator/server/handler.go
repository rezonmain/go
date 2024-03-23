package server

import (
	"fmt"
	"net/http"

	"github.com/rezonmain/cc-validator/cc"
)

func handleValidate() {
	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		cardNumber := r.URL.Query().Get("cn")
		if len(cardNumber) == 0 {
			http.Error(w, "Missing card number, try adding `?cn=xxx` to your request", http.StatusBadRequest)
			return
		}
		isValid := cc.NewCardValidator(cardNumber).ValidateLength().ValidateLuhn().GetIsValid()
		if isValid == false {
			w.Write([]byte("Provided card number is invalid"))
			return
		}
		w.Write([]byte("Provided card number is valid!"))
	})
}

func StartServer() {
	handleValidate()
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
