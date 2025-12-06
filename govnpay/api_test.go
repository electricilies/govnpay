package govnpay_test

import (
	"testing"

	"github.com/electricilies/govnpay/govnpay"
	"github.com/electricilies/govnpay/helper"
	govnpaymodels "github.com/electricilies/govnpay/model"
)

func TestVerifyIPN(t *testing.T) {
	param := &govnpaymodels.VerifyIPNRequest{
		Amount:            "488000000",
		BankTranNo:        "VNP15323638",
		BankCode:          "NCB",
		CardType:          "ATM",
		OrderInfo:         "Payment for order: 019af2cc-7988-772a-af20-6fa5321e9b34",
		PayDate:           "20251206153451",
		ResponseCode:      "00",
		SecureHash:        "93b42d44dd01769a96610efca3429e976bc7ed9542a260a17c47b0d725afbae2",
		TmnCode:           "TNG0ZZLI",
		TransactionNo:     "15323638",
		TransactionStatus: "00",
		TxnRef:            "019af2cc-7988-772a-af20-6fa5321e9b34",

		HashSecret: "4FQY347KK38XDD21RICDD513Z7N91D6C",
		HashAlgo:   helper.Sha256,
	}
	ok, err := govnpay.VerifyIPN(param)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !ok {
		t.Errorf("Expected verification to be successful, but it failed")
	}
}
