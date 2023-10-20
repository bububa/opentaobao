package aesolution

import (
	"sync"
)

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

var poolGlobalAeopTpOperationLogDto = sync.Pool{
	New: func() any {
		return new(GlobalAeopTpOperationLogDto)
	},
}

// GetGlobalAeopTpOperationLogDto() 从对象池中获取GlobalAeopTpOperationLogDto
func GetGlobalAeopTpOperationLogDto() *GlobalAeopTpOperationLogDto {
	return poolGlobalAeopTpOperationLogDto.Get().(*GlobalAeopTpOperationLogDto)
}

// ReleaseGlobalAeopTpOperationLogDto 释放GlobalAeopTpOperationLogDto
func ReleaseGlobalAeopTpOperationLogDto(v *GlobalAeopTpOperationLogDto) {
	v.GmtModified = ""
	v.Memo = ""
	v.ActionType = ""
	v.Operator = ""
	v.GmtCreate = ""
	v.Id = 0
	v.ChildOrderId = 0
	v.OrderId = 0
	poolGlobalAeopTpOperationLogDto.Put(v)
}
