package travel

import (
	"sync"
)

// PontusTravelItemScenic 结构体
type PontusTravelItemScenic struct {
	// 必填，景点所在城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 必填，景点名称
	CnName string `json:"cn_name,omitempty" xml:"cn_name,omitempty"`
	// 景点的经纬度信息，经度在前，纬度在后，英文逗号分隔 最多支持6位小数，如120.111222,30.111222
	Poi string `json:"poi,omitempty" xml:"poi,omitempty"`
	// POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE
	PoiResource string `json:"poi_resource,omitempty" xml:"poi_resource,omitempty"`
	// 必填，门票类型
	TicketType string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
}

var poolPontusTravelItemScenic = sync.Pool{
	New: func() any {
		return new(PontusTravelItemScenic)
	},
}

// GetPontusTravelItemScenic() 从对象池中获取PontusTravelItemScenic
func GetPontusTravelItemScenic() *PontusTravelItemScenic {
	return poolPontusTravelItemScenic.Get().(*PontusTravelItemScenic)
}

// ReleasePontusTravelItemScenic 释放PontusTravelItemScenic
func ReleasePontusTravelItemScenic(v *PontusTravelItemScenic) {
	v.City = ""
	v.CnName = ""
	v.Poi = ""
	v.PoiResource = ""
	v.TicketType = ""
	poolPontusTravelItemScenic.Put(v)
}
