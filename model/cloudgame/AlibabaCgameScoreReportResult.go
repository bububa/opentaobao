package cloudgame

import (
	"sync"
)

// AlibabaCgameScoreReportResult 结构体
type AlibabaCgameScoreReportResult struct {
	// 结果码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果消息
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolAlibabaCgameScoreReportResult = sync.Pool{
	New: func() any {
		return new(AlibabaCgameScoreReportResult)
	},
}

// GetAlibabaCgameScoreReportResult() 从对象池中获取AlibabaCgameScoreReportResult
func GetAlibabaCgameScoreReportResult() *AlibabaCgameScoreReportResult {
	return poolAlibabaCgameScoreReportResult.Get().(*AlibabaCgameScoreReportResult)
}

// ReleaseAlibabaCgameScoreReportResult 释放AlibabaCgameScoreReportResult
func ReleaseAlibabaCgameScoreReportResult(v *AlibabaCgameScoreReportResult) {
	v.Code = ""
	v.Data = ""
	v.Message = ""
	poolAlibabaCgameScoreReportResult.Put(v)
}
