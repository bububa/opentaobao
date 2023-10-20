package hotelalliance

import (
	"sync"
)

// AllianceInfoRequest 结构体
type AllianceInfoRequest struct {
	// 要查询的日期，格式yyyymmdd
	QueryDay string `json:"query_day,omitempty" xml:"query_day,omitempty"`
	// 签约类型-0:融合；1:直签。
	SignType int64 `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
}

var poolAllianceInfoRequest = sync.Pool{
	New: func() any {
		return new(AllianceInfoRequest)
	},
}

// GetAllianceInfoRequest() 从对象池中获取AllianceInfoRequest
func GetAllianceInfoRequest() *AllianceInfoRequest {
	return poolAllianceInfoRequest.Get().(*AllianceInfoRequest)
}

// ReleaseAllianceInfoRequest 释放AllianceInfoRequest
func ReleaseAllianceInfoRequest(v *AllianceInfoRequest) {
	v.QueryDay = ""
	v.SignType = 0
	poolAllianceInfoRequest.Put(v)
}
