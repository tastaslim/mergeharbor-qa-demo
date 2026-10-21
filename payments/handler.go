package payments

import (
	"net/http"
	"strconv"
)

// RefundHandler refunds part of a charge. The amount comes straight from the request.
func RefundHandler(gw Gateway, lookup func(id string) (*Charge, bool)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, ok := lookup(r.FormValue("charge"))
		if !ok {
			http.NotFound(w, r)
			return
		}
		amount, err := strconv.ParseInt(r.FormValue("amount"), 10, 64)
		if err != nil {
			http.Error(w, "invalid amount", http.StatusBadRequest)
			return
		}
		if err := Refund(gw, c, amount); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
