package nrt

import (
	"sync"
)

// TmallNrtPayMerchantStallSigningModifyResultDo 结构体
type TmallNrtPayMerchantStallSigningModifyResultDo struct {
	// 系统自动生成
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 系统自动生成
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 系统自动生成
	Data *StallSigningRespDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallNrtPayMerchantStallSigningModifyResultDo = sync.Pool{
	New: func() any {
		return new(TmallNrtPayMerchantStallSigningModifyResultDo)
	},
}

// GetTmallNrtPayMerchantStallSigningModifyResultDo() 从对象池中获取TmallNrtPayMerchantStallSigningModifyResultDo
func GetTmallNrtPayMerchantStallSigningModifyResultDo() *TmallNrtPayMerchantStallSigningModifyResultDo {
	return poolTmallNrtPayMerchantStallSigningModifyResultDo.Get().(*TmallNrtPayMerchantStallSigningModifyResultDo)
}

// ReleaseTmallNrtPayMerchantStallSigningModifyResultDo 释放TmallNrtPayMerchantStallSigningModifyResultDo
func ReleaseTmallNrtPayMerchantStallSigningModifyResultDo(v *TmallNrtPayMerchantStallSigningModifyResultDo) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Data = nil
	v.Success = false
	poolTmallNrtPayMerchantStallSigningModifyResultDo.Put(v)
}
