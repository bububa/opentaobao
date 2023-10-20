package tuanhotel

import (
	"sync"
)

// TopTuanItemSkuVo 结构体
type TopTuanItemSkuVo struct {
	// 套餐原价，单位为元，仅支持精确到分（小数点后两位）
	OrigPrice string `json:"orig_price,omitempty" xml:"orig_price,omitempty"`
	// 套餐价格。单位为元，仅支持精确到分（小数点后两位）
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 套餐名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 套餐日历库存价格信息，日历库存需完整填写
	CalendarInfo *TopSkuCalendarInfo `json:"calendar_info,omitempty" xml:"calendar_info,omitempty"`
	// skuId，若更新sku信息，必填；若新增sku，此项填写为0
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 套餐间夜
	NightCount int64 `json:"night_count,omitempty" xml:"night_count,omitempty"`
	// 套餐库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 套餐人数
	PeopleCount int64 `json:"people_count,omitempty" xml:"people_count,omitempty"`
	// 使用次数
	UseCount int64 `json:"use_count,omitempty" xml:"use_count,omitempty"`
}

var poolTopTuanItemSkuVo = sync.Pool{
	New: func() any {
		return new(TopTuanItemSkuVo)
	},
}

// GetTopTuanItemSkuVo() 从对象池中获取TopTuanItemSkuVo
func GetTopTuanItemSkuVo() *TopTuanItemSkuVo {
	return poolTopTuanItemSkuVo.Get().(*TopTuanItemSkuVo)
}

// ReleaseTopTuanItemSkuVo 释放TopTuanItemSkuVo
func ReleaseTopTuanItemSkuVo(v *TopTuanItemSkuVo) {
	v.OrigPrice = ""
	v.Price = ""
	v.Name = ""
	v.OuterId = ""
	v.CalendarInfo = nil
	v.SkuId = 0
	v.NightCount = 0
	v.Quantity = 0
	v.PeopleCount = 0
	v.UseCount = 0
	poolTopTuanItemSkuVo.Put(v)
}
