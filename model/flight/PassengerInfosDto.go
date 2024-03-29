package flight

import (
	"sync"
)

// PassengerInfosDto 结构体
type PassengerInfosDto struct {
	// 乘机人基础信息
	PassengerBaseInfos []PassengerBaseInfo `json:"passenger_base_infos,omitempty" xml:"passenger_base_infos>passenger_base_info,omitempty"`
	// 机建（分）
	BuildPrice int64 `json:"build_price,omitempty" xml:"build_price,omitempty"`
	// 数量
	Nums int64 `json:"nums,omitempty" xml:"nums,omitempty"`
	// 燃油（分）
	OilPrice int64 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	// 乘机人类型,1成人2儿童3婴儿
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 销售价
	SalePrice int64 `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 票面价
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

var poolPassengerInfosDto = sync.Pool{
	New: func() any {
		return new(PassengerInfosDto)
	},
}

// GetPassengerInfosDto() 从对象池中获取PassengerInfosDto
func GetPassengerInfosDto() *PassengerInfosDto {
	return poolPassengerInfosDto.Get().(*PassengerInfosDto)
}

// ReleasePassengerInfosDto 释放PassengerInfosDto
func ReleasePassengerInfosDto(v *PassengerInfosDto) {
	v.PassengerBaseInfos = v.PassengerBaseInfos[:0]
	v.BuildPrice = 0
	v.Nums = 0
	v.OilPrice = 0
	v.PassengerType = 0
	v.SalePrice = 0
	v.TicketPrice = 0
	poolPassengerInfosDto.Put(v)
}
