package legalsuit

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
