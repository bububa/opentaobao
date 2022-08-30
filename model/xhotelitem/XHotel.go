package xhotelitem

// XHotel 结构体
type XHotel struct {
	// 卖家自己系统的id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 酒店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 国家编码
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 曾用名
	UsedName string `json:"used_name,omitempty" xml:"used_name,omitempty"`
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
	// 未通过时的拒绝原因等。
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
	CreditCardTypes string `json:"credit_card_types,omitempty" xml:"credit_card_types,omitempty"`
	// 卖家酒店英文名称
	NameE string `json:"name_e,omitempty" xml:"name_e,omitempty"`
	// 对接系统商名称
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 创建时间
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 品牌
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 照片格式，json
	Pics string `json:"pics,omitempty" xml:"pics,omitempty"`
	// 设施
	HotelFacilities string `json:"hotel_facilities,omitempty" xml:"hotel_facilities,omitempty"`
	// 入住政策
	HotelPolicies string `json:"hotel_policies,omitempty" xml:"hotel_policies,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 楼层
	Floors string `json:"floors,omitempty" xml:"floors,omitempty"`
	// 装修时间
	DecorateTime string `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
	// 开业时间
	OpeningTime string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 预订须知
	BookingNotice string `json:"booking_notice,omitempty" xml:"booking_notice,omitempty"`
	// 酒店ID
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 酒店状态：0: 正常;-2:停售；-1：删除
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 淘宝标准酒店信息
	SHotel *SHotel `json:"s_hotel,omitempty" xml:"s_hotel,omitempty"`
	// 0:国内;1:国外
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 省份编码
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 城市编码
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 地区编码
	District int64 `json:"district,omitempty" xml:"district,omitempty"`
	// 已废弃
	MatchStatus int64 `json:"match_status,omitempty" xml:"match_status,omitempty"`
	// 淘宝标准门店信息
	Shotel *XsHotel `json:"shotel,omitempty" xml:"shotel,omitempty"`
	// 房间数
	Rooms int64 `json:"rooms,omitempty" xml:"rooms,omitempty"`
	// 0:酒店；1:客栈
	HotelType int64 `json:"hotel_type,omitempty" xml:"hotel_type,omitempty"`
	// 0:可以接待外宾；1:仅内宾
	ServiceType int64 `json:"service_type,omitempty" xml:"service_type,omitempty"`
}
