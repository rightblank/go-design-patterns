package creational

import "errors"

type PaymentMethod interface {
    Pay(amount float32) float32
}

type PayType int

const (
    Cash  PayType = 1
    Debit PayType = 2
)

func GetPaymentMethod(pt PayType) (PaymentMethod, error) {
    switch pt {
    case Cash:
        return new(CashPayment), nil
    case Debit:
        return new(DebitPayment), nil
    default:
        return nil, errors.New("not implemented")
    }
}

type CashPayment struct {
}

func (*CashPayment) Pay(amount float32) float32 {
    return 2.3
}

type DebitPayment struct {
}

func (*DebitPayment) Pay(amount float32) float32 {
    return 3.2
}
