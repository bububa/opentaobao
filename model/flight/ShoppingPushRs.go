package flight

import (
	"sync"
)

// ShoppingPushRs 结构体
type ShoppingPushRs struct {
	// errRoutingMsg
	ErrRoutingMsgList []string `json:"err_routing_msg_list,omitempty" xml:"err_routing_msg_list>string,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolShoppingPushRs = sync.Pool{
	New: func() any {
		return new(ShoppingPushRs)
	},
}

// GetShoppingPushRs() 从对象池中获取ShoppingPushRs
func GetShoppingPushRs() *ShoppingPushRs {
	return poolShoppingPushRs.Get().(*ShoppingPushRs)
}

// ReleaseShoppingPushRs 释放ShoppingPushRs
func ReleaseShoppingPushRs(v *ShoppingPushRs) {
	v.ErrRoutingMsgList = v.ErrRoutingMsgList[:0]
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolShoppingPushRs.Put(v)
}
