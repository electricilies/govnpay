package govnpaymodels

import (
	"github.com/electricilies/govnpay/helper"
)

type VerifyIPNRequest struct {
	Amount            string
	BankTranNo        string
	BankCode          string
	CardType          string
	OrderInfo         string
	PayDate           string
	ResponseCode      string
	SecureHash        string
	TmnCode           string
	TransactionNo     string
	TransactionStatus string
	TxnRef            string

	HashSecret string
	HashAlgo   helper.HashAlgo
}
