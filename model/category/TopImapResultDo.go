package category

import (
	"sync"
)

// TopImapResultDo 结构体
type TopImapResultDo struct {
	// 返回的pv对列表
	TopPvPairDoList []TopPVPairDo `json:"top_pv_pair_do_list,omitempty" xml:"top_pv_pair_do_list>top_pv_pair_do,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// true表示调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopImapResultDo = sync.Pool{
	New: func() any {
		return new(TopImapResultDo)
	},
}

// GetTopImapResultDo() 从对象池中获取TopImapResultDo
func GetTopImapResultDo() *TopImapResultDo {
	return poolTopImapResultDo.Get().(*TopImapResultDo)
}

// ReleaseTopImapResultDo 释放TopImapResultDo
func ReleaseTopImapResultDo(v *TopImapResultDo) {
	v.TopPvPairDoList = v.TopPvPairDoList[:0]
	v.ErrorMsg = ""
	v.Success = false
	poolTopImapResultDo.Put(v)
}
