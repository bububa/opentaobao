package jym

import (
	"sync"
)

// TaobaoJymMemberRealnameVerifyWithoutuidResultDto 结构体
type TaobaoJymMemberRealnameVerifyWithoutuidResultDto struct {
	// 调用接口结果编码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 调用接口异常信息明细
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 1
	Result *RealNameVerifyTopDto `json:"result,omitempty" xml:"result,omitempty"`
	// 调用是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

var poolTaobaoJymMemberRealnameVerifyWithoutuidResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoJymMemberRealnameVerifyWithoutuidResultDto)
	},
}

// GetTaobaoJymMemberRealnameVerifyWithoutuidResultDto() 从对象池中获取TaobaoJymMemberRealnameVerifyWithoutuidResultDto
func GetTaobaoJymMemberRealnameVerifyWithoutuidResultDto() *TaobaoJymMemberRealnameVerifyWithoutuidResultDto {
	return poolTaobaoJymMemberRealnameVerifyWithoutuidResultDto.Get().(*TaobaoJymMemberRealnameVerifyWithoutuidResultDto)
}

// ReleaseTaobaoJymMemberRealnameVerifyWithoutuidResultDto 释放TaobaoJymMemberRealnameVerifyWithoutuidResultDto
func ReleaseTaobaoJymMemberRealnameVerifyWithoutuidResultDto(v *TaobaoJymMemberRealnameVerifyWithoutuidResultDto) {
	v.StateCode = ""
	v.ExtraErrMsg = ""
	v.Result = nil
	v.Succ = false
	poolTaobaoJymMemberRealnameVerifyWithoutuidResultDto.Put(v)
}
