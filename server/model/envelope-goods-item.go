package model

import "github.com/shopspring/decimal"

type EnvelopeGoodsItem struct {
	BaseModel
	ItemNo       string            `json:"itemNo"`
	EnvelopeNo   string            `json:"envelopNo"`
	RecvUsername string            `json:"recvUsername"`
	RecvUserId   int64             `json:"recvUserId"`
	Amount       decimal.Decimal   `json:"amount"`       // 收到的金额
	RemainAmount decimal.Decimal   `json:"remainAmount"` // 剩余金额
	AccountNo    string            `json:"accountNo"`
	PayStatus    EnvelopePayStatus `json:"payStatus"`
	Desc         string            `json:"desc"`
	RecvUser     User              `json:"recvUser" gorm:"foreignKey:RecvUserId"`
}

type RedEnvelopeReceiveDTO struct {
	EnvelopeNo   string `json:"envelopeNo" validate:"required"` // 红包编号,红包唯一标识
	RecvUserId   int64  `json:"recvUserId"`                     // 红包接收者用户编号
	RecvUsername string `json:"recvUsername"`                   // 红包接收者用户名称
	AccountNo    string `json:"accountNo"`
}
