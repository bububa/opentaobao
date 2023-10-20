package drugtrace

import (
	"sync"
)

// CodeToBill 结构体
type CodeToBill struct {
	// codeToBill列表
	BillIdentityList []BillIdentity `json:"bill_identity_list,omitempty" xml:"bill_identity_list>bill_identity,omitempty"`
	// 追溯码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolCodeToBill = sync.Pool{
	New: func() any {
		return new(CodeToBill)
	},
}

// GetCodeToBill() 从对象池中获取CodeToBill
func GetCodeToBill() *CodeToBill {
	return poolCodeToBill.Get().(*CodeToBill)
}

// ReleaseCodeToBill 释放CodeToBill
func ReleaseCodeToBill(v *CodeToBill) {
	v.BillIdentityList = v.BillIdentityList[:0]
	v.Code = ""
	poolCodeToBill.Put(v)
}
