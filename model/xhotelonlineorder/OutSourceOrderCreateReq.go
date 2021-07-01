package xhotelonlineorder

// OutSourceOrderCreateReq 结构体
type OutSourceOrderCreateReq struct {
	// 酒店hid
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 房间数
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 下单来源信息
	CreateAppKey string `json:"create_app_key,omitempty" xml:"create_app_key,omitempty"`
	// 是否即时确认订单,1表示是
	TagJsqr int64 `json:"tag_jsqr,omitempty" xml:"tag_jsqr,omitempty"`
	// gid
	Gid int64 `json:"gid,omitempty" xml:"gid,omitempty"`
	// 备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// rid
	Rid int64 `json:"rid,omitempty" xml:"rid,omitempty"`
	// 预订的标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// rpid
	RpId int64 `json:"rp_id,omitempty" xml:"rp_id,omitempty"`
	// 外部渠道用户id，钉钉必传
	OutUserId string `json:"out_user_id,omitempty" xml:"out_user_id,omitempty"`
	// rate id
	RateId int64 `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	// srid
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 宝贝id
	ItemIid int64 `json:"item_iid,omitempty" xml:"item_iid,omitempty"`
	// 外部渠道 钉钉 欢行 ...
	OutSource string `json:"out_source,omitempty" xml:"out_source,omitempty"`
	// 是否直连订单
	IsDirectOrder bool `json:"is_direct_order,omitempty" xml:"is_direct_order,omitempty"`
	// 最晚到店时间
	LateArriveTime string `json:"late_arrive_time,omitempty" xml:"late_arrive_time,omitempty"`
	// 离店日期
	Checkout string `json:"checkout,omitempty" xml:"checkout,omitempty"`
	// 联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 房型名称
	RoomTypeName string `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// ttid信息
	Ttid string `json:"ttid,omitempty" xml:"ttid,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 入住人信息
	GuestInfoList []GuestInfo `json:"guest_info_list,omitempty" xml:"guest_info_list>guest_info,omitempty"`
	// 入住日期
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 外部订单号，幂等性保障
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 总价  单位 分
	TotalRoomPrice int64 `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
	// 日历价格
	DailyPriceInfoList []DailyPriceInfo `json:"daily_price_info_list,omitempty" xml:"daily_price_info_list>daily_price_info,omitempty"`
	// 一些扩展属性
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 取消政策
	CancelPolicy *CancelPolicy `json:"cancel_policy,omitempty" xml:"cancel_policy,omitempty"`
	// 库存类型,0 rate普通库存 1 room普通库存 2 rate协议保留房 3 rate普通保留房
	InventoryTypeFlag int64 `json:"inventory_type_flag,omitempty" xml:"inventory_type_flag,omitempty"`
	// 取消政策描述
	CancelPolicyDesc string `json:"cancel_policy_desc,omitempty" xml:"cancel_policy_desc,omitempty"`
	// 平台促销
	PlatformPromotion int64 `json:"platform_promotion,omitempty" xml:"platform_promotion,omitempty"`
	// 卖家促销
	SellerPromotion int64 `json:"seller_promotion,omitempty" xml:"seller_promotion,omitempty"`
}
