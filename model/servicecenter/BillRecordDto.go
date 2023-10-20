package servicecenter

import (
	"sync"
)

// BillRecordDto 结构体
type BillRecordDto struct {
	// appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 备用字段
	Extend1 string `json:"extend1,omitempty" xml:"extend1,omitempty"`
	// 备用字段
	Extend10 string `json:"extend10,omitempty" xml:"extend10,omitempty"`
	// 备用字段
	Extend2 string `json:"extend2,omitempty" xml:"extend2,omitempty"`
	// 备用字段
	Extend3 string `json:"extend3,omitempty" xml:"extend3,omitempty"`
	// 备用字段
	Extend4 string `json:"extend4,omitempty" xml:"extend4,omitempty"`
	// 备用字段
	Extend5 string `json:"extend5,omitempty" xml:"extend5,omitempty"`
	// 备用字段
	Extend6 string `json:"extend6,omitempty" xml:"extend6,omitempty"`
	// 备用字段
	Extend7 string `json:"extend7,omitempty" xml:"extend7,omitempty"`
	// 备用字段
	Extend8 string `json:"extend8,omitempty" xml:"extend8,omitempty"`
	// 备用字段
	Extend9 string `json:"extend9,omitempty" xml:"extend9,omitempty"`
	// 卖家nick
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 外部确认账单ID
	OutConfirmId string `json:"out_confirm_id,omitempty" xml:"out_confirm_id,omitempty"`
	// 外部订单ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 记录产生时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 目标号码
	TargetNo string `json:"target_no,omitempty" xml:"target_no,omitempty"`
	// 金额
	Fee int64 `json:"fee,omitempty" xml:"fee,omitempty"`
	// 服务市场有效订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 状态：1成功、2失败
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 账单分类：1短信
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolBillRecordDto = sync.Pool{
	New: func() any {
		return new(BillRecordDto)
	},
}

// GetBillRecordDto() 从对象池中获取BillRecordDto
func GetBillRecordDto() *BillRecordDto {
	return poolBillRecordDto.Get().(*BillRecordDto)
}

// ReleaseBillRecordDto 释放BillRecordDto
func ReleaseBillRecordDto(v *BillRecordDto) {
	v.Appkey = ""
	v.Extend1 = ""
	v.Extend10 = ""
	v.Extend2 = ""
	v.Extend3 = ""
	v.Extend4 = ""
	v.Extend5 = ""
	v.Extend6 = ""
	v.Extend7 = ""
	v.Extend8 = ""
	v.Extend9 = ""
	v.Nick = ""
	v.OutConfirmId = ""
	v.OutOrderId = ""
	v.StartDate = ""
	v.TargetNo = ""
	v.Fee = 0
	v.OrderId = 0
	v.Status = 0
	v.Type = 0
	poolBillRecordDto.Put(v)
}
