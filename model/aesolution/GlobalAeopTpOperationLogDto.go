package aesolution

// GlobalAeopTpOperationLogDto 结构体
type GlobalAeopTpOperationLogDto struct {
	// order modification time
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// buyer memo
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// action type
	ActionType string `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// operator
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// order creation time
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// child order ID
	ChildOrderId int64 `json:"child_order_id,omitempty" xml:"child_order_id,omitempty"`
	// order ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
