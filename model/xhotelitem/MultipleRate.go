package xhotelitem

import (
	"sync"
)

// MultipleRate 结构体
type MultipleRate struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 创建时间
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 价格和库存信息,包括加床价，加人价等信息。date  日期必须为 T---T+90 日内的日期（T为当天），且不能重复price 价格 int类型 取值范围1-99999999 单位为分quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)addPerson 加人价addBed 加床价
	InventoryPrice string `json:"inventory_price,omitempty" xml:"inventory_price,omitempty"`
	// 入住人数
	Occupancy int64 `json:"occupancy,omitempty" xml:"occupancy,omitempty"`
	// 连住天数
	Lengthofstay int64 `json:"lengthofstay,omitempty" xml:"lengthofstay,omitempty"`
	// 酒店商品id
	Gid int64 `json:"gid,omitempty" xml:"gid,omitempty"`
	// 房价id
	Rpid int64 `json:"rpid,omitempty" xml:"rpid,omitempty"`
	// 币种
	CurrencyCode int64 `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
}

var poolMultipleRate = sync.Pool{
	New: func() any {
		return new(MultipleRate)
	},
}

// GetMultipleRate() 从对象池中获取MultipleRate
func GetMultipleRate() *MultipleRate {
	return poolMultipleRate.Get().(*MultipleRate)
}

// ReleaseMultipleRate 释放MultipleRate
func ReleaseMultipleRate(v *MultipleRate) {
	v.Name = ""
	v.CreatedTime = ""
	v.ModifiedTime = ""
	v.InventoryPrice = ""
	v.Occupancy = 0
	v.Lengthofstay = 0
	v.Gid = 0
	v.Rpid = 0
	v.CurrencyCode = 0
	poolMultipleRate.Put(v)
}
