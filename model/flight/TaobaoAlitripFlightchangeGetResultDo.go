package flight

import (
	"sync"
)

// TaobaoAlitripFlightchangeGetResultDo 结构体
type TaobaoAlitripFlightchangeGetResultDo struct {
	// results
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripFlightchangeGetResultDo = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripFlightchangeGetResultDo)
	},
}

// GetTaobaoAlitripFlightchangeGetResultDo() 从对象池中获取TaobaoAlitripFlightchangeGetResultDo
func GetTaobaoAlitripFlightchangeGetResultDo() *TaobaoAlitripFlightchangeGetResultDo {
	return poolTaobaoAlitripFlightchangeGetResultDo.Get().(*TaobaoAlitripFlightchangeGetResultDo)
}

// ReleaseTaobaoAlitripFlightchangeGetResultDo 释放TaobaoAlitripFlightchangeGetResultDo
func ReleaseTaobaoAlitripFlightchangeGetResultDo(v *TaobaoAlitripFlightchangeGetResultDo) {
	v.Results = v.Results[:0]
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolTaobaoAlitripFlightchangeGetResultDo.Put(v)
}
