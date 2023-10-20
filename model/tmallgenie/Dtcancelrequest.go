package tmallgenie

import (
	"sync"
)

// Dtcancelrequest 结构体
type Dtcancelrequest struct {
	// 服务单号
	Zfwdh string `json:"zfwdh,omitempty" xml:"zfwdh,omitempty"`
}

var poolDtcancelrequest = sync.Pool{
	New: func() any {
		return new(Dtcancelrequest)
	},
}

// GetDtcancelrequest() 从对象池中获取Dtcancelrequest
func GetDtcancelrequest() *Dtcancelrequest {
	return poolDtcancelrequest.Get().(*Dtcancelrequest)
}

// ReleaseDtcancelrequest 释放Dtcancelrequest
func ReleaseDtcancelrequest(v *Dtcancelrequest) {
	v.Zfwdh = ""
	poolDtcancelrequest.Put(v)
}
