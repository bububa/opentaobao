package util

import (
	"sync"
)

// BaseRep 结构体
type BaseRep struct {
	// 内层大对象
	Datas []AssetQrCodeDto `json:"datas,omitempty" xml:"datas>asset_qr_code_dto,omitempty"`
	// 返回错误消息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回错误code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolBaseRep = sync.Pool{
	New: func() any {
		return new(BaseRep)
	},
}

// GetBaseRep() 从对象池中获取BaseRep
func GetBaseRep() *BaseRep {
	return poolBaseRep.Get().(*BaseRep)
}

// ReleaseBaseRep 释放BaseRep
func ReleaseBaseRep(v *BaseRep) {
	v.Datas = v.Datas[:0]
	v.Msg = ""
	v.Code = ""
	v.Data = false
	poolBaseRep.Put(v)
}
