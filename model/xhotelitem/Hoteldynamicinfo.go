package xhotelitem

// Hoteldynamicinfo 结构体
type Hoteldynamicinfo struct {
	// 不可售原因
	UnsaleReason string `json:"unsale_reason,omitempty" xml:"unsale_reason,omitempty"`
	// 酒店hid
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 酒店状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 酒店的销售渠道
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 可售健康房型数
	KsHeathyRoomNum int64 `json:"ks_heathy_room_num,omitempty" xml:"ks_heathy_room_num,omitempty"`
	// 电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 标准酒店ID
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 城市名称
	CityStr string `json:"city_str,omitempty" xml:"city_str,omitempty"`
	// 城市编码
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 房型数
	RoomNun int64 `json:"room_nun,omitempty" xml:"room_nun,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 可售房型数
	KsRoomNum int64 `json:"ks_room_num,omitempty" xml:"ks_room_num,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 不可售原因
	UnsaleType int64 `json:"unsale_type,omitempty" xml:"unsale_type,omitempty"`
	// 酒店名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 分析日期
	CalculateDate string `json:"calculate_date,omitempty" xml:"calculate_date,omitempty"`
	// 酒店匹配
	DataConfirm int64 `json:"data_confirm,omitempty" xml:"data_confirm,omitempty"`
	// 酒店外部ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
}
