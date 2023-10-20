package cloudgame

import (
	"sync"
)

// AlibabaCgameMpMpprojectInitnewprojectResult 结构体
type AlibabaCgameMpMpprojectInitnewprojectResult struct {
	// 0
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// sucess
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// login session
	Data *MpProjectConfigDto `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCgameMpMpprojectInitnewprojectResult = sync.Pool{
	New: func() any {
		return new(AlibabaCgameMpMpprojectInitnewprojectResult)
	},
}

// GetAlibabaCgameMpMpprojectInitnewprojectResult() 从对象池中获取AlibabaCgameMpMpprojectInitnewprojectResult
func GetAlibabaCgameMpMpprojectInitnewprojectResult() *AlibabaCgameMpMpprojectInitnewprojectResult {
	return poolAlibabaCgameMpMpprojectInitnewprojectResult.Get().(*AlibabaCgameMpMpprojectInitnewprojectResult)
}

// ReleaseAlibabaCgameMpMpprojectInitnewprojectResult 释放AlibabaCgameMpMpprojectInitnewprojectResult
func ReleaseAlibabaCgameMpMpprojectInitnewprojectResult(v *AlibabaCgameMpMpprojectInitnewprojectResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCgameMpMpprojectInitnewprojectResult.Put(v)
}
