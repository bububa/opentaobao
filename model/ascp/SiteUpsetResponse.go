package ascp

import (
	"sync"
)

// SiteUpsetResponse 结构体
type SiteUpsetResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolSiteUpsetResponse = sync.Pool{
	New: func() any {
		return new(SiteUpsetResponse)
	},
}

// GetSiteUpsetResponse() 从对象池中获取SiteUpsetResponse
func GetSiteUpsetResponse() *SiteUpsetResponse {
	return poolSiteUpsetResponse.Get().(*SiteUpsetResponse)
}

// ReleaseSiteUpsetResponse 释放SiteUpsetResponse
func ReleaseSiteUpsetResponse(v *SiteUpsetResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolSiteUpsetResponse.Put(v)
}
