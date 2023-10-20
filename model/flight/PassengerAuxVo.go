package flight

import (
	"sync"
)

// PassengerAuxVo 结构体
type PassengerAuxVo struct {
	// 乘机人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 辅营产品规格信息
	AuxProductVo *AuxProductVo `json:"aux_product_vo,omitempty" xml:"aux_product_vo,omitempty"`
	// 购买数量
	BookNum int64 `json:"book_num,omitempty" xml:"book_num,omitempty"`
	// 航段信息
	FlightVo *BookFlightVo `json:"flight_vo,omitempty" xml:"flight_vo,omitempty"`
}

var poolPassengerAuxVo = sync.Pool{
	New: func() any {
		return new(PassengerAuxVo)
	},
}

// GetPassengerAuxVo() 从对象池中获取PassengerAuxVo
func GetPassengerAuxVo() *PassengerAuxVo {
	return poolPassengerAuxVo.Get().(*PassengerAuxVo)
}

// ReleasePassengerAuxVo 释放PassengerAuxVo
func ReleasePassengerAuxVo(v *PassengerAuxVo) {
	v.Name = ""
	v.AuxProductVo = nil
	v.BookNum = 0
	v.FlightVo = nil
	poolPassengerAuxVo.Put(v)
}
