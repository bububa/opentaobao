package cloudgame

import (
	"sync"
)

// AlibabaCgameMpMpprojectLoginexistaccountResult 结构体
type AlibabaCgameMpMpprojectLoginexistaccountResult struct {
	// 0
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// sucess
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// login session
	Data *LoginSessionDto `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaCgameMpMpprojectLoginexistaccountResult = sync.Pool{
	New: func() any {
		return new(AlibabaCgameMpMpprojectLoginexistaccountResult)
	},
}

// GetAlibabaCgameMpMpprojectLoginexistaccountResult() 从对象池中获取AlibabaCgameMpMpprojectLoginexistaccountResult
func GetAlibabaCgameMpMpprojectLoginexistaccountResult() *AlibabaCgameMpMpprojectLoginexistaccountResult {
	return poolAlibabaCgameMpMpprojectLoginexistaccountResult.Get().(*AlibabaCgameMpMpprojectLoginexistaccountResult)
}

// ReleaseAlibabaCgameMpMpprojectLoginexistaccountResult 释放AlibabaCgameMpMpprojectLoginexistaccountResult
func ReleaseAlibabaCgameMpMpprojectLoginexistaccountResult(v *AlibabaCgameMpMpprojectLoginexistaccountResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolAlibabaCgameMpMpprojectLoginexistaccountResult.Put(v)
}
