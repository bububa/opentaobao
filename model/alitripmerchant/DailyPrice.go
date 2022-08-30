package alitripmerchant

// DailyPrice 结构体
type DailyPrice struct {
	// 每日价格
	ReallyPrice string `json:"really_price,omitempty" xml:"really_price,omitempty"`
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 外币每日价格
	OutReallyPrice string `json:"out_really_price,omitempty" xml:"out_really_price,omitempty"`
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 床型图标
	BedTypeIcon string `json:"bed_type_icon,omitempty" xml:"bed_type_icon,omitempty"`
	// 入住人数图标
	MaxOccupancyIcon string `json:"max_occupancy_icon,omitempty" xml:"max_occupancy_icon,omitempty"`
	// 楼层
	Floor string `json:"floor,omitempty" xml:"floor,omitempty"`
	// 楼层图标
	FloorIcon string `json:"floor_icon,omitempty" xml:"floor_icon,omitempty"`
	// 面积图标
	AreaIcon string `json:"area_icon,omitempty" xml:"area_icon,omitempty"`
	// 床型图标
	WindowTypeIcon string `json:"window_type_icon,omitempty" xml:"window_type_icon,omitempty"`
	// 窗型
	WindowType string `json:"window_type,omitempty" xml:"window_type,omitempty"`
	// 最大入住人数
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
}
