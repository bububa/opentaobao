package hotel

// RoomTypeDailyPrice 结构体
type RoomTypeDailyPrice struct {
	// 库价具体信息
	Rates []Rate `json:"rates,omitempty" xml:"rates>rate,omitempty"`
	// 床型数据
	BedTypeString string `json:"bed_type_string,omitempty" xml:"bed_type_string,omitempty"`
	// h5下单页链接
	H5BuyUrl string `json:"h5_buy_url,omitempty" xml:"h5_buy_url,omitempty"`
	// 房型名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// pc下单页链接
	PcBuyUrl string `json:"pc_buy_url,omitempty" xml:"pc_buy_url,omitempty"`
	// 床型信息
	WindowType string `json:"window_type,omitempty" xml:"window_type,omitempty"`
	// 卖家房型id
	Rid int64 `json:"rid,omitempty" xml:"rid,omitempty"`
	// 标准房型srid
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
}
