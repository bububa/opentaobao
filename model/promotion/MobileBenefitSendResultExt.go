package promotion

import (
	"sync"
)

// MobileBenefitSendResultExt 结构体
type MobileBenefitSendResultExt struct {
	// 错误码列表
	ErrorCodeList []string `json:"error_code_list,omitempty" xml:"error_code_list>string,omitempty"`
	// 失败份数
	FailureCount int64 `json:"failure_count,omitempty" xml:"failure_count,omitempty"`
	// 成功份数
	SuccessCount int64 `json:"success_count,omitempty" xml:"success_count,omitempty"`
	// 活动详情id
	IndexId int64 `json:"index_id,omitempty" xml:"index_id,omitempty"`
}

var poolMobileBenefitSendResultExt = sync.Pool{
	New: func() any {
		return new(MobileBenefitSendResultExt)
	},
}

// GetMobileBenefitSendResultExt() 从对象池中获取MobileBenefitSendResultExt
func GetMobileBenefitSendResultExt() *MobileBenefitSendResultExt {
	return poolMobileBenefitSendResultExt.Get().(*MobileBenefitSendResultExt)
}

// ReleaseMobileBenefitSendResultExt 释放MobileBenefitSendResultExt
func ReleaseMobileBenefitSendResultExt(v *MobileBenefitSendResultExt) {
	v.ErrorCodeList = v.ErrorCodeList[:0]
	v.FailureCount = 0
	v.SuccessCount = 0
	v.IndexId = 0
	poolMobileBenefitSendResultExt.Put(v)
}
