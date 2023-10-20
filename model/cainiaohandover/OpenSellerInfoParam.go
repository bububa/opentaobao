package cainiaohandover

import (
	"sync"
)

// OpenSellerInfoParam 结构体
type OpenSellerInfoParam struct {
	// 跨店铺组包时的店铺分组ID
	TopUserKey string `json:"top_user_key,omitempty" xml:"top_user_key,omitempty"`
}

var poolOpenSellerInfoParam = sync.Pool{
	New: func() any {
		return new(OpenSellerInfoParam)
	},
}

// GetOpenSellerInfoParam() 从对象池中获取OpenSellerInfoParam
func GetOpenSellerInfoParam() *OpenSellerInfoParam {
	return poolOpenSellerInfoParam.Get().(*OpenSellerInfoParam)
}

// ReleaseOpenSellerInfoParam 释放OpenSellerInfoParam
func ReleaseOpenSellerInfoParam(v *OpenSellerInfoParam) {
	v.TopUserKey = ""
	poolOpenSellerInfoParam.Put(v)
}
