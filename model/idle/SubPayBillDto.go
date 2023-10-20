package idle

import (
	"sync"
)

// SubPayBillDto 结构体
type SubPayBillDto struct {
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 支付状态描述
	PayStatusDesc string `json:"pay_status_desc,omitempty" xml:"pay_status_desc,omitempty"`
	// 支付状态
	PayStatus string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 代扣错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 代扣错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 支付宝流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 计划id
	PlanId string `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
}

var poolSubPayBillDto = sync.Pool{
	New: func() any {
		return new(SubPayBillDto)
	},
}

// GetSubPayBillDto() 从对象池中获取SubPayBillDto
func GetSubPayBillDto() *SubPayBillDto {
	return poolSubPayBillDto.Get().(*SubPayBillDto)
}

// ReleaseSubPayBillDto 释放SubPayBillDto
func ReleaseSubPayBillDto(v *SubPayBillDto) {
	v.PayTime = ""
	v.CreateTime = ""
	v.PayStatusDesc = ""
	v.PayStatus = ""
	v.ErrMsg = ""
	v.ErrCode = ""
	v.AlipayTradeNo = ""
	v.Amount = ""
	v.PlanId = ""
	poolSubPayBillDto.Put(v)
}
