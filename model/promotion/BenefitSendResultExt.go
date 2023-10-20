package promotion

import (
	"sync"
)

// BenefitSendResultExt 结构体
type BenefitSendResultExt struct {
	// 异常码列表
	ErrorCodeList []string `json:"error_code_list,omitempty" xml:"error_code_list>string,omitempty"`
	// 失败数
	FailureCount int64 `json:"failure_count,omitempty" xml:"failure_count,omitempty"`
	// 活动详情id
	IndexId int64 `json:"index_id,omitempty" xml:"index_id,omitempty"`
	// 成功数
	SuccessCount int64 `json:"success_count,omitempty" xml:"success_count,omitempty"`
}

var poolBenefitSendResultExt = sync.Pool{
	New: func() any {
		return new(BenefitSendResultExt)
	},
}

// GetBenefitSendResultExt() 从对象池中获取BenefitSendResultExt
func GetBenefitSendResultExt() *BenefitSendResultExt {
	return poolBenefitSendResultExt.Get().(*BenefitSendResultExt)
}

// ReleaseBenefitSendResultExt 释放BenefitSendResultExt
func ReleaseBenefitSendResultExt(v *BenefitSendResultExt) {
	v.ErrorCodeList = v.ErrorCodeList[:0]
	v.FailureCount = 0
	v.IndexId = 0
	v.SuccessCount = 0
	poolBenefitSendResultExt.Put(v)
}
