package cloudgame

import (
	"sync"
)

// HavanaUserIdQueryRequest 结构体
type HavanaUserIdQueryRequest struct {
	// API名称
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 前端APPKEY
	FrontAppKey string `json:"front_app_key,omitempty" xml:"front_app_key,omitempty"`
	// 登录态token
	AccountToken string `json:"account_token,omitempty" xml:"account_token,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolHavanaUserIdQueryRequest = sync.Pool{
	New: func() any {
		return new(HavanaUserIdQueryRequest)
	},
}

// GetHavanaUserIdQueryRequest() 从对象池中获取HavanaUserIdQueryRequest
func GetHavanaUserIdQueryRequest() *HavanaUserIdQueryRequest {
	return poolHavanaUserIdQueryRequest.Get().(*HavanaUserIdQueryRequest)
}

// ReleaseHavanaUserIdQueryRequest 释放HavanaUserIdQueryRequest
func ReleaseHavanaUserIdQueryRequest(v *HavanaUserIdQueryRequest) {
	v.Action = ""
	v.FrontAppKey = ""
	v.AccountToken = ""
	v.Version = 0
	poolHavanaUserIdQueryRequest.Put(v)
}
