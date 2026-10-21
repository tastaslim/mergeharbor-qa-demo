package payments

import "testing"

type fakeGW struct{ calls int }

func (f *fakeGW) Refund(string, int64) error { f.calls++; return nil }

func TestRefundWithinCharge(t *testing.T) {
	gw := &fakeGW{}
	c := &Charge{ID: "ch_1", Amount: 1000}
	if err := Refund(gw, c, 400); err != nil || c.Refunded != 400 || gw.calls != 1 {
		t.Fatalf("err=%v refunded=%d calls=%d", err, c.Refunded, gw.calls)
	}
}
