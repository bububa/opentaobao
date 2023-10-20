package legalsuit

import (
	"sync"
)

// PaymentOrderModel 结构体
type PaymentOrderModel struct {
	// 附件信息
	PaymentOrderFiles []FileModel `json:"payment_order_files,omitempty" xml:"payment_order_files>file_model,omitempty"`
	// 推送人
	SendPeople string `json:"send_people,omitempty" xml:"send_people,omitempty"`
	// 案号
	CaseNumber string `json:"case_number,omitempty" xml:"case_number,omitempty"`
	// 备注
	Comment string `json:"comment,omitempty" xml:"comment,omitempty"`
	// 案件id
	SuitId int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
	// 委托id
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
}

var poolPaymentOrderModel = sync.Pool{
	New: func() any {
		return new(PaymentOrderModel)
	},
}

// GetPaymentOrderModel() 从对象池中获取PaymentOrderModel
func GetPaymentOrderModel() *PaymentOrderModel {
	return poolPaymentOrderModel.Get().(*PaymentOrderModel)
}

// ReleasePaymentOrderModel 释放PaymentOrderModel
func ReleasePaymentOrderModel(v *PaymentOrderModel) {
	v.PaymentOrderFiles = v.PaymentOrderFiles[:0]
	v.SendPeople = ""
	v.CaseNumber = ""
	v.Comment = ""
	v.SuitId = 0
	v.EntrustId = 0
	poolPaymentOrderModel.Put(v)
}
