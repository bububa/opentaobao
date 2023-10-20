package c2m

import (
	"sync"
)

// TaobaoSebpOrganizationGetorderinfoResultDo 结构体
type TaobaoSebpOrganizationGetorderinfoResultDo struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 结果信息
	Module *PageInfo `json:"module,omitempty" xml:"module,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoSebpOrganizationGetorderinfoResultDo = sync.Pool{
	New: func() any {
		return new(TaobaoSebpOrganizationGetorderinfoResultDo)
	},
}

// GetTaobaoSebpOrganizationGetorderinfoResultDo() 从对象池中获取TaobaoSebpOrganizationGetorderinfoResultDo
func GetTaobaoSebpOrganizationGetorderinfoResultDo() *TaobaoSebpOrganizationGetorderinfoResultDo {
	return poolTaobaoSebpOrganizationGetorderinfoResultDo.Get().(*TaobaoSebpOrganizationGetorderinfoResultDo)
}

// ReleaseTaobaoSebpOrganizationGetorderinfoResultDo 释放TaobaoSebpOrganizationGetorderinfoResultDo
func ReleaseTaobaoSebpOrganizationGetorderinfoResultDo(v *TaobaoSebpOrganizationGetorderinfoResultDo) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Module = nil
	v.Success = false
	poolTaobaoSebpOrganizationGetorderinfoResultDo.Put(v)
}
