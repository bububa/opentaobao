package c2m

import (
	"sync"
)

// TaobaoSebpOrganizationGetinviteinfoResultDo 结构体
type TaobaoSebpOrganizationGetinviteinfoResultDo struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 结果信息
	Module *PageInfo `json:"module,omitempty" xml:"module,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoSebpOrganizationGetinviteinfoResultDo = sync.Pool{
	New: func() any {
		return new(TaobaoSebpOrganizationGetinviteinfoResultDo)
	},
}

// GetTaobaoSebpOrganizationGetinviteinfoResultDo() 从对象池中获取TaobaoSebpOrganizationGetinviteinfoResultDo
func GetTaobaoSebpOrganizationGetinviteinfoResultDo() *TaobaoSebpOrganizationGetinviteinfoResultDo {
	return poolTaobaoSebpOrganizationGetinviteinfoResultDo.Get().(*TaobaoSebpOrganizationGetinviteinfoResultDo)
}

// ReleaseTaobaoSebpOrganizationGetinviteinfoResultDo 释放TaobaoSebpOrganizationGetinviteinfoResultDo
func ReleaseTaobaoSebpOrganizationGetinviteinfoResultDo(v *TaobaoSebpOrganizationGetinviteinfoResultDo) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Module = nil
	v.Success = false
	poolTaobaoSebpOrganizationGetinviteinfoResultDo.Put(v)
}
