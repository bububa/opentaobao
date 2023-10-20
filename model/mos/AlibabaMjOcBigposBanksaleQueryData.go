package mos

import (
	"sync"
)

// AlibabaMjOcBigposBanksaleQueryData 结构体
type AlibabaMjOcBigposBanksaleQueryData struct {
	// 交易时间
	OperTime string `json:"oper_time,omitempty" xml:"oper_time,omitempty"`
	// 查询流水号
	PosTraceNo string `json:"pos_trace_no,omitempty" xml:"pos_trace_no,omitempty"`
	// 银联系统参考号
	Refnum string `json:"refnum,omitempty" xml:"refnum,omitempty"`
	// 小票号后7位
	Fphm string `json:"fphm,omitempty" xml:"fphm,omitempty"`
	// 行号
	RowNo int64 `json:"row_no,omitempty" xml:"row_no,omitempty"`
	// 交易金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 已调账金额
	AdjustedAmount int64 `json:"adjusted_amount,omitempty" xml:"adjusted_amount,omitempty"`
}

var poolAlibabaMjOcBigposBanksaleQueryData = sync.Pool{
	New: func() any {
		return new(AlibabaMjOcBigposBanksaleQueryData)
	},
}

// GetAlibabaMjOcBigposBanksaleQueryData() 从对象池中获取AlibabaMjOcBigposBanksaleQueryData
func GetAlibabaMjOcBigposBanksaleQueryData() *AlibabaMjOcBigposBanksaleQueryData {
	return poolAlibabaMjOcBigposBanksaleQueryData.Get().(*AlibabaMjOcBigposBanksaleQueryData)
}

// ReleaseAlibabaMjOcBigposBanksaleQueryData 释放AlibabaMjOcBigposBanksaleQueryData
func ReleaseAlibabaMjOcBigposBanksaleQueryData(v *AlibabaMjOcBigposBanksaleQueryData) {
	v.OperTime = ""
	v.PosTraceNo = ""
	v.Refnum = ""
	v.Fphm = ""
	v.RowNo = 0
	v.Amount = 0
	v.AdjustedAmount = 0
	poolAlibabaMjOcBigposBanksaleQueryData.Put(v)
}
