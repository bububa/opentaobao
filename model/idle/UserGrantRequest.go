package idle

import (
	"sync"
)

// UserGrantRequest 结构体
type UserGrantRequest struct {
	// 当前用户的所属业务类型编码，优品&amp;开放平台业务 默认使用 IDLE_TOP
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 场景码，标识品类。22:虚拟卡券/账号
	SceneType string `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
}

var poolUserGrantRequest = sync.Pool{
	New: func() any {
		return new(UserGrantRequest)
	},
}

// GetUserGrantRequest() 从对象池中获取UserGrantRequest
func GetUserGrantRequest() *UserGrantRequest {
	return poolUserGrantRequest.Get().(*UserGrantRequest)
}

// ReleaseUserGrantRequest 释放UserGrantRequest
func ReleaseUserGrantRequest(v *UserGrantRequest) {
	v.BizCode = ""
	v.SceneType = ""
	poolUserGrantRequest.Put(v)
}
