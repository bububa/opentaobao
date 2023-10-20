package tbtrade

import (
	"sync"
)

// StepPayDetail 结构体
type StepPayDetail struct {
	// 支付宝交易号
	StepChannelNo string `json:"step_channel_no,omitempty" xml:"step_channel_no,omitempty"`
	// 支付方式
	StepInstrumentCode string `json:"step_instrument_code,omitempty" xml:"step_instrument_code,omitempty"`
	// 支付金额
	StepActualPayFee string `json:"step_actual_pay_fee,omitempty" xml:"step_actual_pay_fee,omitempty"`
	// 分阶段支付单Id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 支付状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 支付顺序
	StepNo int64 `json:"step_no,omitempty" xml:"step_no,omitempty"`
}

var poolStepPayDetail = sync.Pool{
	New: func() any {
		return new(StepPayDetail)
	},
}

// GetStepPayDetail() 从对象池中获取StepPayDetail
func GetStepPayDetail() *StepPayDetail {
	return poolStepPayDetail.Get().(*StepPayDetail)
}

// ReleaseStepPayDetail 释放StepPayDetail
func ReleaseStepPayDetail(v *StepPayDetail) {
	v.StepChannelNo = ""
	v.StepInstrumentCode = ""
	v.StepActualPayFee = ""
	v.Id = 0
	v.PayStatus = 0
	v.StepNo = 0
	poolStepPayDetail.Put(v)
}
