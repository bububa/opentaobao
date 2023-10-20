package tbk

import (
	"sync"
)

// TaobaoTbkDgVegasTljCreateResult 结构体
type TaobaoTbkDgVegasTljCreateResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *RightsInstanceCreateResult `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTbkDgVegasTljCreateResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgVegasTljCreateResult)
	},
}

// GetTaobaoTbkDgVegasTljCreateResult() 从对象池中获取TaobaoTbkDgVegasTljCreateResult
func GetTaobaoTbkDgVegasTljCreateResult() *TaobaoTbkDgVegasTljCreateResult {
	return poolTaobaoTbkDgVegasTljCreateResult.Get().(*TaobaoTbkDgVegasTljCreateResult)
}

// ReleaseTaobaoTbkDgVegasTljCreateResult 释放TaobaoTbkDgVegasTljCreateResult
func ReleaseTaobaoTbkDgVegasTljCreateResult(v *TaobaoTbkDgVegasTljCreateResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolTaobaoTbkDgVegasTljCreateResult.Put(v)
}
