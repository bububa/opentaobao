package servicecenter

import (
	"sync"
)

// IncomeConfirmDto 结构体
type IncomeConfirmDto struct {
	// appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 确认扩展信息
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 卖家nick
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 外部确认账单ID
	OutConfirmId string `json:"out_confirm_id,omitempty" xml:"out_confirm_id,omitempty"`
	// 外部订单ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 确认金额
	Fee int64 `json:"fee,omitempty" xml:"fee,omitempty"`
	// 服务市场有效订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolIncomeConfirmDto = sync.Pool{
	New: func() any {
		return new(IncomeConfirmDto)
	},
}

// GetIncomeConfirmDto() 从对象池中获取IncomeConfirmDto
func GetIncomeConfirmDto() *IncomeConfirmDto {
	return poolIncomeConfirmDto.Get().(*IncomeConfirmDto)
}

// ReleaseIncomeConfirmDto 释放IncomeConfirmDto
func ReleaseIncomeConfirmDto(v *IncomeConfirmDto) {
	v.Appkey = ""
	v.Extend = ""
	v.Nick = ""
	v.OutConfirmId = ""
	v.OutOrderId = ""
	v.Fee = 0
	v.OrderId = 0
	poolIncomeConfirmDto.Put(v)
}
