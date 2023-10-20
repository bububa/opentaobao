package waybill

import (
	"sync"
)

// CainiaoWaybillPrivacySubscriptionGetResult 结构体
type CainiaoWaybillPrivacySubscriptionGetResult struct {
	// 错误code列表
	ErrorCodeList []string `json:"error_code_list,omitempty" xml:"error_code_list>string,omitempty"`
	// 错误列表
	ErrorInfoList []string `json:"error_info_list,omitempty" xml:"error_info_list>string,omitempty"`
	// 第一个错误
	OneErrorInfo string `json:"one_error_info,omitempty" xml:"one_error_info,omitempty"`
	// 系统自动生成
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 系统信息
	ObjectId string `json:"object_id,omitempty" xml:"object_id,omitempty"`
	// 是否失败
	Failure bool `json:"failure,omitempty" xml:"failure,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 商家是否订购
	Subscription bool `json:"subscription,omitempty" xml:"subscription,omitempty"`
}

var poolCainiaoWaybillPrivacySubscriptionGetResult = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillPrivacySubscriptionGetResult)
	},
}

// GetCainiaoWaybillPrivacySubscriptionGetResult() 从对象池中获取CainiaoWaybillPrivacySubscriptionGetResult
func GetCainiaoWaybillPrivacySubscriptionGetResult() *CainiaoWaybillPrivacySubscriptionGetResult {
	return poolCainiaoWaybillPrivacySubscriptionGetResult.Get().(*CainiaoWaybillPrivacySubscriptionGetResult)
}

// ReleaseCainiaoWaybillPrivacySubscriptionGetResult 释放CainiaoWaybillPrivacySubscriptionGetResult
func ReleaseCainiaoWaybillPrivacySubscriptionGetResult(v *CainiaoWaybillPrivacySubscriptionGetResult) {
	v.ErrorCodeList = v.ErrorCodeList[:0]
	v.ErrorInfoList = v.ErrorInfoList[:0]
	v.OneErrorInfo = ""
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.ObjectId = ""
	v.Failure = false
	v.Success = false
	v.Subscription = false
	poolCainiaoWaybillPrivacySubscriptionGetResult.Put(v)
}
