package user

import (
	"sync"
)

// TaobaoAlitripFlightchangeAddResultDo 结构体
type TaobaoAlitripFlightchangeAddResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripFlightchangeAddResultDo = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripFlightchangeAddResultDo)
	},
}

// GetTaobaoAlitripFlightchangeAddResultDo() 从对象池中获取TaobaoAlitripFlightchangeAddResultDo
func GetTaobaoAlitripFlightchangeAddResultDo() *TaobaoAlitripFlightchangeAddResultDo {
	return poolTaobaoAlitripFlightchangeAddResultDo.Get().(*TaobaoAlitripFlightchangeAddResultDo)
}

// ReleaseTaobaoAlitripFlightchangeAddResultDo 释放TaobaoAlitripFlightchangeAddResultDo
func ReleaseTaobaoAlitripFlightchangeAddResultDo(v *TaobaoAlitripFlightchangeAddResultDo) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolTaobaoAlitripFlightchangeAddResultDo.Put(v)
}
