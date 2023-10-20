package tbk

import (
	"sync"
)

// TaobaoTbkScInvitecodeGetData 结构体
type TaobaoTbkScInvitecodeGetData struct {
	// 邀请码
	InviterCode string `json:"inviter_code,omitempty" xml:"inviter_code,omitempty"`
}

var poolTaobaoTbkScInvitecodeGetData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScInvitecodeGetData)
	},
}

// GetTaobaoTbkScInvitecodeGetData() 从对象池中获取TaobaoTbkScInvitecodeGetData
func GetTaobaoTbkScInvitecodeGetData() *TaobaoTbkScInvitecodeGetData {
	return poolTaobaoTbkScInvitecodeGetData.Get().(*TaobaoTbkScInvitecodeGetData)
}

// ReleaseTaobaoTbkScInvitecodeGetData 释放TaobaoTbkScInvitecodeGetData
func ReleaseTaobaoTbkScInvitecodeGetData(v *TaobaoTbkScInvitecodeGetData) {
	v.InviterCode = ""
	poolTaobaoTbkScInvitecodeGetData.Put(v)
}
