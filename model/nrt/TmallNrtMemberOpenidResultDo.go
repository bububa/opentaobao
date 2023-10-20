package nrt

import (
	"sync"
)

// TmallNrtMemberOpenidResultDo 结构体
type TmallNrtMemberOpenidResultDo struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回的业务数据
	Data *MemberSynResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

var poolTmallNrtMemberOpenidResultDo = sync.Pool{
	New: func() any {
		return new(TmallNrtMemberOpenidResultDo)
	},
}

// GetTmallNrtMemberOpenidResultDo() 从对象池中获取TmallNrtMemberOpenidResultDo
func GetTmallNrtMemberOpenidResultDo() *TmallNrtMemberOpenidResultDo {
	return poolTmallNrtMemberOpenidResultDo.Get().(*TmallNrtMemberOpenidResultDo)
}

// ReleaseTmallNrtMemberOpenidResultDo 释放TmallNrtMemberOpenidResultDo
func ReleaseTmallNrtMemberOpenidResultDo(v *TmallNrtMemberOpenidResultDo) {
	v.Code = ""
	v.Msg = ""
	v.Data = nil
	v.Succ = false
	poolTmallNrtMemberOpenidResultDo.Put(v)
}
