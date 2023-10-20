package cloudgame

import (
	"sync"
)

// TopQueryUserBodyDressRequest 结构体
type TopQueryUserBodyDressRequest struct {
	// 云游戏MixUserId
	MixUserId string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
	// API版本号, 可选
	Version string `json:"version,omitempty" xml:"version,omitempty"`
}

var poolTopQueryUserBodyDressRequest = sync.Pool{
	New: func() any {
		return new(TopQueryUserBodyDressRequest)
	},
}

// GetTopQueryUserBodyDressRequest() 从对象池中获取TopQueryUserBodyDressRequest
func GetTopQueryUserBodyDressRequest() *TopQueryUserBodyDressRequest {
	return poolTopQueryUserBodyDressRequest.Get().(*TopQueryUserBodyDressRequest)
}

// ReleaseTopQueryUserBodyDressRequest 释放TopQueryUserBodyDressRequest
func ReleaseTopQueryUserBodyDressRequest(v *TopQueryUserBodyDressRequest) {
	v.MixUserId = ""
	v.Version = ""
	poolTopQueryUserBodyDressRequest.Put(v)
}
