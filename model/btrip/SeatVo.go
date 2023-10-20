package btrip

import (
	"sync"
)

// SeatVo 结构体
type SeatVo struct {
	// 名称
	SeatName string `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
	// 候补价
	HoubuPrice int64 `json:"houbu_price,omitempty" xml:"houbu_price,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 类型
	SeatType int64 `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// 下铺价
	SleeperPrice int64 `json:"sleeper_price,omitempty" xml:"sleeper_price,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
}

var poolSeatVo = sync.Pool{
	New: func() any {
		return new(SeatVo)
	},
}

// GetSeatVo() 从对象池中获取SeatVo
func GetSeatVo() *SeatVo {
	return poolSeatVo.Get().(*SeatVo)
}

// ReleaseSeatVo 释放SeatVo
func ReleaseSeatVo(v *SeatVo) {
	v.SeatName = ""
	v.HoubuPrice = 0
	v.Price = 0
	v.SeatType = 0
	v.SleeperPrice = 0
	v.Stock = 0
	poolSeatVo.Put(v)
}
