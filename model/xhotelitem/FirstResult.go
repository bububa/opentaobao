package xhotelitem

// FirstResult 结构体
type FirstResult struct {
	// 未通过时的拒绝原因等。
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// 创建时间
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 卖家自己系统的id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 酒店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 曾用名
	UsedName string `json:"used_name,omitempty" xml:"used_name,omitempty"`
	// 国家编码
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 商圈信息
	Business string `json:"business,omitempty" xml:"business,omitempty"`
	// 酒店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 坐标类型
	PositionType string `json:"position_type,omitempty" xml:"position_type,omitempty"`
	// 酒店电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 扩展信息
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 匹配结果
	DataConfirmStr string `json:"data_confirm_str,omitempty" xml:"data_confirm_str,omitempty"`
	// 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
	CreditCardTypes string `json:"credit_card_types,omitempty" xml:"credit_card_types,omitempty"`
	// 卖家酒店英文名称
	NameE string `json:"name_e,omitempty" xml:"name_e,omitempty"`
	// vendor
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 货币类型（编码,字母编码）,hid 维度支持的币种信息,目前只能 add 时添加，不支持 update时更新,,如果DB中是null ，则默认是人民币 CNY
	CurrencyCodeName string `json:"currency_code_name,omitempty" xml:"currency_code_name,omitempty"`
	// 酒店维度特殊标签含义, json: {"pure-direct-hotel":0,"direct-manual-order-hotel":1,"ebk-direct-hotel":0,"non-direct-hotel":1,"allow-dingding-business-travel-hotel":1,"willing-dingding-bussiness-travel-hotel":0,"calendar-room-package-hotel":1,"dijiajiajia-hotel":0,"ebk-number-of-confirm-room-hotel":1} , key含义: pure-direct-hotel 表示纯直连酒店, direct-manual-order-hotel 和 ebk-direct-hotel 和 non-direct-hotel 这三个key对应value都是0 . direct-manual-order-hotel 表示 人工承接失败订单的酒店标签。如果某个直连酒店打了该标签，那么直连下单失败以后，允许人工承接订单，由人工跟进 . ebk-direct-hotel 表示 ebk直连酒店标。如果某个酒店打了该标签，那么这个酒店下允许通过ebk发布直连rp . non-direct-hotel 表示 卖家非直连酒店标签。如果某个酒店打了该标签，那么该酒店下单不会走直连交易。 allow-dingding-business-travel-hotel 表示 允许进入阿里商旅渠道（钉钉）售卖信用住的单体酒店 willing-dingding-bussiness-travel-hotel 表示 已签协议愿意加入阿里商旅渠道售卖信用住的单体酒店 . calendar-room-package-hotel 表示 酒店可以参加日历房套餐活动包括创建，修改，删除活动信息（高星集团GMV项目） dijiajiajia-hotel 表示 底价加价酒店权限标。只有打了该标的酒店才允许维护底价加价规则和包房rp . ebk-number-of-confirm-room-hotel 表示ebk确认订单，是否要输入外部确认号 . nonstandard-project-hotel 表示该酒店是否参加非标项目
	TagJson string `json:"tag_json,omitempty" xml:"tag_json,omitempty"`
	// 酒店对应的旺旺号
	AliNick string `json:"ali_nick,omitempty" xml:"ali_nick,omitempty"`
	// 资源方房型设施
	StandardRoomFacilities string `json:"standard_room_facilities,omitempty" xml:"standard_room_facilities,omitempty"`
	// 资源方酒店服务
	StandardHotelService string `json:"standard_hotel_service,omitempty" xml:"standard_hotel_service,omitempty"`
	// 资源方酒店设施
	StandardHotelFacilities string `json:"standard_hotel_facilities,omitempty" xml:"standard_hotel_facilities,omitempty"`
	// 资源方预订须知
	StandardBookingNotice string `json:"standard_booking_notice,omitempty" xml:"standard_booking_notice,omitempty"`
	// 资源方娱乐设施
	StandardAmuseFacilities string `json:"standard_amuse_facilities,omitempty" xml:"standard_amuse_facilities,omitempty"`
	// 商品下架原因
	DownReason string `json:"down_reason,omitempty" xml:"down_reason,omitempty"`
	// out_rid
	OutRid string `json:"out_rid,omitempty" xml:"out_rid,omitempty"`
	// 宝贝名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 购买须知
	Guide string `json:"guide,omitempty" xml:"guide,omitempty"`
	// 宝贝描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 酒店商品图片Url。多个url用逗号隔开
	PicUrls string `json:"pic_urls,omitempty" xml:"pic_urls,omitempty"`
	// 发票类型。A,B。分别代表： A:酒店住宿发票,B:其他
	ReceiptType string `json:"receipt_type,omitempty" xml:"receipt_type,omitempty"`
	// 发票类型为其他时的发票描述,不能超过30个汉字，60个字符
	ReceiptOtherTypeDesc string `json:"receipt_other_type_desc,omitempty" xml:"receipt_other_type_desc,omitempty"`
	// 发票说明，不能超过100个汉字,200个字符。
	ReceiptInfo string `json:"receipt_info,omitempty" xml:"receipt_info,omitempty"`
	// 库存日历
	Inventory string `json:"inventory,omitempty" xml:"inventory,omitempty"`
	// extend_info1
	ExtendInfo1 string `json:"extend_info1,omitempty" xml:"extend_info1,omitempty"`
	// extend_info2
	ExtendInfo2 string `json:"extend_info2,omitempty" xml:"extend_info2,omitempty"`
	// extend_info3
	ExtendInfo3 string `json:"extend_info3,omitempty" xml:"extend_info3,omitempty"`
	// switchCalendar
	SwitchCalendar string `json:"switch_calendar,omitempty" xml:"switch_calendar,omitempty"`
	// 酒店ID
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 酒店状态：0: 正常;-2:停售；-1：删除
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 淘宝标准酒店信息
	SHotel *SHotel `json:"s_hotel,omitempty" xml:"s_hotel,omitempty"`
	// hotel匹配状态: 0：待系统匹配 1：已系统匹配，匹配成功，待卖家确认 2：已系统匹配，匹配失败，待人工匹配 3：已人工匹配，匹配成功，待卖家确认 4：已人工匹配，匹配失败 5：卖家已确认，确认&ldquo;YES&rdquo; 6：卖家已确认，确认&ldquo;NO&rdquo; 7:已系统匹配，但是匹配重复，待人工确认
	MatchStatus int64 `json:"match_status,omitempty" xml:"match_status,omitempty"`
	// 0:国内;1:国外
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 省份编码
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 城市编码
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 地区编码
	District int64 `json:"district,omitempty" xml:"district,omitempty"`
	// 标识该酒店所走的 结算流程，如果是null 默认是 国内结算流程，否则是其他的，比如：海外信用住结算流程
	BillingProcessType int64 `json:"billing_process_type,omitempty" xml:"billing_process_type,omitempty"`
	// 离线数据,该酒店在售1,不在售0,未知-1
	OnSale int64 `json:"on_sale,omitempty" xml:"on_sale,omitempty"`
	// 离线数据,该酒店热搜1,非热搜0,未知-1
	HotSearch int64 `json:"hot_search,omitempty" xml:"hot_search,omitempty"`
	// 离线数据,该酒店热卖1,非热卖0,未知-1
	HotSale int64 `json:"hot_sale,omitempty" xml:"hot_sale,omitempty"`
	// gid酒店商品id
	Gid int64 `json:"gid,omitempty" xml:"gid,omitempty"`
	// iid淘宝商品id
	Iid int64 `json:"iid,omitempty" xml:"iid,omitempty"`
	// rid房型id
	Rid int64 `json:"rid,omitempty" xml:"rid,omitempty"`
	// 酒店商品是否提供发票
	HasReceipt bool `json:"has_receipt,omitempty" xml:"has_receipt,omitempty"`
	// 橱窗推荐
	Recommend bool `json:"recommend,omitempty" xml:"recommend,omitempty"`
}
