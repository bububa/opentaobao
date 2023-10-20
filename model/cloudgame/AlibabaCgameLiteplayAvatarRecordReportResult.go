package cloudgame

import (
	"sync"
)

// AlibabaCgameLiteplayAvatarRecordReportResult 结构体
type AlibabaCgameLiteplayAvatarRecordReportResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// str消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回消息体
	Data *TopRecordCallbackResp `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCgameLiteplayAvatarRecordReportResult = sync.Pool{
	New: func() any {
		return new(AlibabaCgameLiteplayAvatarRecordReportResult)
	},
}

// GetAlibabaCgameLiteplayAvatarRecordReportResult() 从对象池中获取AlibabaCgameLiteplayAvatarRecordReportResult
func GetAlibabaCgameLiteplayAvatarRecordReportResult() *AlibabaCgameLiteplayAvatarRecordReportResult {
	return poolAlibabaCgameLiteplayAvatarRecordReportResult.Get().(*AlibabaCgameLiteplayAvatarRecordReportResult)
}

// ReleaseAlibabaCgameLiteplayAvatarRecordReportResult 释放AlibabaCgameLiteplayAvatarRecordReportResult
func ReleaseAlibabaCgameLiteplayAvatarRecordReportResult(v *AlibabaCgameLiteplayAvatarRecordReportResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCgameLiteplayAvatarRecordReportResult.Put(v)
}
