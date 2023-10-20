package simba

import (
	"sync"
)

// AdzoneConfigVo 结构体
type AdzoneConfigVo struct {
	// 描述
	AimDesc string `json:"aim_desc,omitempty" xml:"aim_desc,omitempty"`
	// 资源包id
	AdzoneId int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
	// 是否支持溢价
	Discount bool `json:"discount,omitempty" xml:"discount,omitempty"`
}

var poolAdzoneConfigVo = sync.Pool{
	New: func() any {
		return new(AdzoneConfigVo)
	},
}

// GetAdzoneConfigVo() 从对象池中获取AdzoneConfigVo
func GetAdzoneConfigVo() *AdzoneConfigVo {
	return poolAdzoneConfigVo.Get().(*AdzoneConfigVo)
}

// ReleaseAdzoneConfigVo 释放AdzoneConfigVo
func ReleaseAdzoneConfigVo(v *AdzoneConfigVo) {
	v.AimDesc = ""
	v.AdzoneId = 0
	v.Discount = false
	poolAdzoneConfigVo.Put(v)
}
