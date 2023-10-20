package tmallgenie

import (
	"sync"
)

// Dtcancelresponse 结构体
type Dtcancelresponse struct {
	// 返回消息
	ReturnMessage *Dtreturnmessage `json:"return_message,omitempty" xml:"return_message,omitempty"`
}

var poolDtcancelresponse = sync.Pool{
	New: func() any {
		return new(Dtcancelresponse)
	},
}

// GetDtcancelresponse() 从对象池中获取Dtcancelresponse
func GetDtcancelresponse() *Dtcancelresponse {
	return poolDtcancelresponse.Get().(*Dtcancelresponse)
}

// ReleaseDtcancelresponse 释放Dtcancelresponse
func ReleaseDtcancelresponse(v *Dtcancelresponse) {
	v.ReturnMessage = nil
	poolDtcancelresponse.Put(v)
}
