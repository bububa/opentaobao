package alitripmerchant

import (
	"sync"
)

// QueryMemberParam 结构体
type QueryMemberParam struct {
	// 签名
	Signature string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 手机号前缀（加密）
	PhonePre string `json:"phone_pre,omitempty" xml:"phone_pre,omitempty"`
	// 手机号（加密）
	PhoneNum string `json:"phone_num,omitempty" xml:"phone_num,omitempty"`
	// 应用id
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
}

var poolQueryMemberParam = sync.Pool{
	New: func() any {
		return new(QueryMemberParam)
	},
}

// GetQueryMemberParam() 从对象池中获取QueryMemberParam
func GetQueryMemberParam() *QueryMemberParam {
	return poolQueryMemberParam.Get().(*QueryMemberParam)
}

// ReleaseQueryMemberParam 释放QueryMemberParam
func ReleaseQueryMemberParam(v *QueryMemberParam) {
	v.Signature = ""
	v.PhonePre = ""
	v.PhoneNum = ""
	v.AppId = ""
	poolQueryMemberParam.Put(v)
}
