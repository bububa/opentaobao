package xhotelonlineorder

import (
	"sync"
)

// OtherFeeDetail 结构体
type OtherFeeDetail struct {
	// 无
	OtherFeeInfos []OtherFeeInfo `json:"other_fee_infos,omitempty" xml:"other_fee_infos>other_fee_info,omitempty"`
}

var poolOtherFeeDetail = sync.Pool{
	New: func() any {
		return new(OtherFeeDetail)
	},
}

// GetOtherFeeDetail() 从对象池中获取OtherFeeDetail
func GetOtherFeeDetail() *OtherFeeDetail {
	return poolOtherFeeDetail.Get().(*OtherFeeDetail)
}

// ReleaseOtherFeeDetail 释放OtherFeeDetail
func ReleaseOtherFeeDetail(v *OtherFeeDetail) {
	v.OtherFeeInfos = v.OtherFeeInfos[:0]
	poolOtherFeeDetail.Put(v)
}
