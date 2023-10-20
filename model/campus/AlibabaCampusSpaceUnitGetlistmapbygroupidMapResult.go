package campus

import (
	"sync"
)

// AlibabaCampusSpaceUnitGetlistmapbygroupidMapResult 结构体
type AlibabaCampusSpaceUnitGetlistmapbygroupidMapResult struct {
	// content
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorLevel
	ErrorLevel string `json:"error_level,omitempty" xml:"error_level,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaCampusSpaceUnitGetlistmapbygroupidMapResult = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceUnitGetlistmapbygroupidMapResult)
	},
}

// GetAlibabaCampusSpaceUnitGetlistmapbygroupidMapResult() 从对象池中获取AlibabaCampusSpaceUnitGetlistmapbygroupidMapResult
func GetAlibabaCampusSpaceUnitGetlistmapbygroupidMapResult() *AlibabaCampusSpaceUnitGetlistmapbygroupidMapResult {
	return poolAlibabaCampusSpaceUnitGetlistmapbygroupidMapResult.Get().(*AlibabaCampusSpaceUnitGetlistmapbygroupidMapResult)
}

// ReleaseAlibabaCampusSpaceUnitGetlistmapbygroupidMapResult 释放AlibabaCampusSpaceUnitGetlistmapbygroupidMapResult
func ReleaseAlibabaCampusSpaceUnitGetlistmapbygroupidMapResult(v *AlibabaCampusSpaceUnitGetlistmapbygroupidMapResult) {
	v.Content = ""
	v.RequestId = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.ErrorLevel = ""
	v.Success = false
	poolAlibabaCampusSpaceUnitGetlistmapbygroupidMapResult.Put(v)
}
