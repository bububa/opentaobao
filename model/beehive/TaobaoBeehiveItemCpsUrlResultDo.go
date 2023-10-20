package beehive

import (
	"sync"
)

// TaobaoBeehiveItemCpsUrlResultDo 结构体
type TaobaoBeehiveItemCpsUrlResultDo struct {
	// 商品id和对应的url map
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoBeehiveItemCpsUrlResultDo = sync.Pool{
	New: func() any {
		return new(TaobaoBeehiveItemCpsUrlResultDo)
	},
}

// GetTaobaoBeehiveItemCpsUrlResultDo() 从对象池中获取TaobaoBeehiveItemCpsUrlResultDo
func GetTaobaoBeehiveItemCpsUrlResultDo() *TaobaoBeehiveItemCpsUrlResultDo {
	return poolTaobaoBeehiveItemCpsUrlResultDo.Get().(*TaobaoBeehiveItemCpsUrlResultDo)
}

// ReleaseTaobaoBeehiveItemCpsUrlResultDo 释放TaobaoBeehiveItemCpsUrlResultDo
func ReleaseTaobaoBeehiveItemCpsUrlResultDo(v *TaobaoBeehiveItemCpsUrlResultDo) {
	v.Model = ""
	v.Code = ""
	v.Msg = ""
	v.Success = false
	poolTaobaoBeehiveItemCpsUrlResultDo.Put(v)
}
