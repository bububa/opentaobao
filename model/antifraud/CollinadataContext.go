package antifraud

import (
	"sync"
)

// CollinadataContext 结构体
type CollinadataContext struct {
	// 客户标识,由ISP指定
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 页面生成的请求的标识.
	TokenId string `json:"token_id,omitempty" xml:"token_id,omitempty"`
	// 场景标识, 由ISP指
	SceneId string `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	// 参数签名, 签名函数由ISP提供
	SerialNo string `json:"serial_no,omitempty" xml:"serial_no,omitempty"`
	// 透传参数
	Trans string `json:"trans,omitempty" xml:"trans,omitempty"`
	// 发起查询的时间,用于加密serial_no
	TimeStamp int64 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
}

var poolCollinadataContext = sync.Pool{
	New: func() any {
		return new(CollinadataContext)
	},
}

// GetCollinadataContext() 从对象池中获取CollinadataContext
func GetCollinadataContext() *CollinadataContext {
	return poolCollinadataContext.Get().(*CollinadataContext)
}

// ReleaseCollinadataContext 释放CollinadataContext
func ReleaseCollinadataContext(v *CollinadataContext) {
	v.AppKey = ""
	v.TokenId = ""
	v.SceneId = ""
	v.SerialNo = ""
	v.Trans = ""
	v.TimeStamp = 0
	poolCollinadataContext.Put(v)
}
