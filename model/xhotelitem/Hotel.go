package xhotelitem

import (
	"sync"
)

// Hotel 结构体
type Hotel struct {
	// 酒店修改备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 酒店设施。json格式
	HotelFacilities string `json:"hotel_facilities,omitempty" xml:"hotel_facilities,omitempty"`
	// 酒店类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 扩展信息
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 邮编
	PostalCode string `json:"postal_code,omitempty" xml:"postal_code,omitempty"`
	// 楼层信息
	Floors string `json:"floors,omitempty" xml:"floors,omitempty"`
	// 卖家名称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 酒店中文描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 货币类型（编码,字母编码）
	CurrencyCodeName string `json:"currency_code_name,omitempty" xml:"currency_code_name,omitempty"`
	// 酒店英文描述
	EnDesc string `json:"en_desc,omitempty" xml:"en_desc,omitempty"`
	// domestic=0时，固定China； domestic=1时，是海外国家编码值
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 酒店入住政策
	HotelPolicies string `json:"hotel_policies,omitempty" xml:"hotel_policies,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 操作人信息
	OperXiaoerName string `json:"oper_xiaoer_name,omitempty" xml:"oper_xiaoer_name,omitempty"`
	// 酒店外部ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 酒店修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 酒店英文地址
	EnAddr string `json:"en_addr,omitempty" xml:"en_addr,omitempty"`
	// 商圈
	Business string `json:"business,omitempty" xml:"business,omitempty"`
	// 酒店曾用名
	UsedName string `json:"used_name,omitempty" xml:"used_name,omitempty"`
	// 酒店图片信息
	Pics string `json:"pics,omitempty" xml:"pics,omitempty"`
	// 房间设施
	RoomFacilities string `json:"room_facilities,omitempty" xml:"room_facilities,omitempty"`
	// 酒店名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 坐标类型，现在支持：G : Google:B : 百度;A : 高德;M : Mapbar;L : 灵图
	PositionType string `json:"position_type,omitempty" xml:"position_type,omitempty"`
	// 酒店名称(英文)
	NameE string `json:"name_e,omitempty" xml:"name_e,omitempty"`
	// 酒店创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 酒店的销售渠道
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 扩展信息
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 酒店星级
	Star string `json:"star,omitempty" xml:"star,omitempty"`
	// 预订须知
	BookingNotice string `json:"booking_notice,omitempty" xml:"booking_notice,omitempty"`
	// 装修时间
	DecorateTime string `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
	// 酒店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 酒店服务
	Service string `json:"service,omitempty" xml:"service,omitempty"`
	// 品牌
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 开业时间
	OpeningTime string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 酒店ID
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 城市编码
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 省份编码
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 匹配是否人工确认
	DataConfirm int64 `json:"data_confirm,omitempty" xml:"data_confirm,omitempty"`
	// 房间数
	Rooms int64 `json:"rooms,omitempty" xml:"rooms,omitempty"`
	// 酒店状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 酒店下架类型
	DownShelfType int64 `json:"down_shelf_type,omitempty" xml:"down_shelf_type,omitempty"`
	// 标准酒店ID
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 酒店支付结算类型
	BillingProcessType int64 `json:"billing_process_type,omitempty" xml:"billing_process_type,omitempty"`
	// 地区编码
	District int64 `json:"district,omitempty" xml:"district,omitempty"`
	// 是否国外。
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 来源
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 匹配状态
	MatchStatus int64 `json:"match_status,omitempty" xml:"match_status,omitempty"`
	// 判断该酒店是不是对应的卖家直营
	KzzyTag int64 `json:"kzzy_tag,omitempty" xml:"kzzy_tag,omitempty"`
}

var poolHotel = sync.Pool{
	New: func() any {
		return new(Hotel)
	},
}

// GetHotel() 从对象池中获取Hotel
func GetHotel() *Hotel {
	return poolHotel.Get().(*Hotel)
}

// ReleaseHotel 释放Hotel
func ReleaseHotel(v *Hotel) {
	v.Remark = ""
	v.Tel = ""
	v.HotelFacilities = ""
	v.Type = ""
	v.Ext = ""
	v.PostalCode = ""
	v.Floors = ""
	v.SellerNick = ""
	v.Description = ""
	v.Longitude = ""
	v.CurrencyCodeName = ""
	v.EnDesc = ""
	v.Country = ""
	v.HotelPolicies = ""
	v.Latitude = ""
	v.OperXiaoerName = ""
	v.OuterId = ""
	v.GmtModified = ""
	v.EnAddr = ""
	v.Business = ""
	v.UsedName = ""
	v.Pics = ""
	v.RoomFacilities = ""
	v.Name = ""
	v.PositionType = ""
	v.NameE = ""
	v.GmtCreate = ""
	v.Vendor = ""
	v.Extend = ""
	v.Star = ""
	v.BookingNotice = ""
	v.DecorateTime = ""
	v.Address = ""
	v.Service = ""
	v.Brand = ""
	v.OpeningTime = ""
	v.Hid = 0
	v.City = 0
	v.Province = 0
	v.DataConfirm = 0
	v.Rooms = 0
	v.Status = 0
	v.DownShelfType = 0
	v.Shid = 0
	v.BillingProcessType = 0
	v.District = 0
	v.Domestic = 0
	v.Source = 0
	v.SellerId = 0
	v.MatchStatus = 0
	v.KzzyTag = 0
	poolHotel.Put(v)
}
