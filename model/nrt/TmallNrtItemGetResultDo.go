package nrt

import (
	"sync"
)

// TmallNrtItemGetResultDo 结构体
type TmallNrtItemGetResultDo struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Data *TopHomeItemDto `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallNrtItemGetResultDo = sync.Pool{
	New: func() any {
		return new(TmallNrtItemGetResultDo)
	},
}

// GetTmallNrtItemGetResultDo() 从对象池中获取TmallNrtItemGetResultDo
func GetTmallNrtItemGetResultDo() *TmallNrtItemGetResultDo {
	return poolTmallNrtItemGetResultDo.Get().(*TmallNrtItemGetResultDo)
}

// ReleaseTmallNrtItemGetResultDo 释放TmallNrtItemGetResultDo
func ReleaseTmallNrtItemGetResultDo(v *TmallNrtItemGetResultDo) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = nil
	v.Success = false
	poolTmallNrtItemGetResultDo.Put(v)
}
