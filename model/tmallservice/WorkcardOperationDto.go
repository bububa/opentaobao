package tmallservice

import (
	"sync"
)

// WorkcardOperationDto 结构体
type WorkcardOperationDto struct {
	// 修改时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 操作属性扩展字段
	AttributeMap string `json:"attribute_map,omitempty" xml:"attribute_map,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 动作描述
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 操作人类型，1:买家,2:服务商,3:网点
	OperatorType int64 `json:"operator_type,omitempty" xml:"operator_type,omitempty"`
}

var poolWorkcardOperationDto = sync.Pool{
	New: func() any {
		return new(WorkcardOperationDto)
	},
}

// GetWorkcardOperationDto() 从对象池中获取WorkcardOperationDto
func GetWorkcardOperationDto() *WorkcardOperationDto {
	return poolWorkcardOperationDto.Get().(*WorkcardOperationDto)
}

// ReleaseWorkcardOperationDto 释放WorkcardOperationDto
func ReleaseWorkcardOperationDto(v *WorkcardOperationDto) {
	v.GmtModify = ""
	v.AttributeMap = ""
	v.GmtCreate = ""
	v.Operator = ""
	v.Action = ""
	v.OperatorType = 0
	poolWorkcardOperationDto.Put(v)
}
