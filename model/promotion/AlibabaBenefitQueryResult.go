package promotion

import (
	"sync"
)

// AlibabaBenefitQueryResult 结构体
type AlibabaBenefitQueryResult struct {
	// datas
	Datas []OrightDto `json:"datas,omitempty" xml:"datas>oright_dto,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaBenefitQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaBenefitQueryResult)
	},
}

// GetAlibabaBenefitQueryResult() 从对象池中获取AlibabaBenefitQueryResult
func GetAlibabaBenefitQueryResult() *AlibabaBenefitQueryResult {
	return poolAlibabaBenefitQueryResult.Get().(*AlibabaBenefitQueryResult)
}

// ReleaseAlibabaBenefitQueryResult 释放AlibabaBenefitQueryResult
func ReleaseAlibabaBenefitQueryResult(v *AlibabaBenefitQueryResult) {
	v.Datas = v.Datas[:0]
	v.Msg = ""
	v.Code = ""
	v.Success = false
	poolAlibabaBenefitQueryResult.Put(v)
}
