package hotel

import (
	"sync"
)

// HotelPriceResultSet 结构体
type HotelPriceResultSet struct {
	// 每个标准酒店的库价集合
	Results []SHotelPrice `json:"results,omitempty" xml:"results>s_hotel_price,omitempty"`
	// 当前用户的会员信息
	BindMemberInfos []SellerSupplierPartnerMemberInfoVo `json:"bind_member_infos,omitempty" xml:"bind_member_infos>seller_supplier_partner_member_info_vo,omitempty"`
	// 渠道id，0--pc,1--无线
	ChannelId string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// errorCode，error=true时有值
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息，error=true时有值
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 报价list页h5链接
	H5ListUrl string `json:"h5_list_url,omitempty" xml:"h5_list_url,omitempty"`
	// list页pc链接
	HotelListUrl string `json:"hotel_list_url,omitempty" xml:"hotel_list_url,omitempty"`
	// 请求id，用于链路追踪
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据总量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 版本信息
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 本次请求是否错误，true--出粗
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
	// 当前数据集合是否还有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// 服务调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolHotelPriceResultSet = sync.Pool{
	New: func() any {
		return new(HotelPriceResultSet)
	},
}

// GetHotelPriceResultSet() 从对象池中获取HotelPriceResultSet
func GetHotelPriceResultSet() *HotelPriceResultSet {
	return poolHotelPriceResultSet.Get().(*HotelPriceResultSet)
}

// ReleaseHotelPriceResultSet 释放HotelPriceResultSet
func ReleaseHotelPriceResultSet(v *HotelPriceResultSet) {
	v.Results = v.Results[:0]
	v.BindMemberInfos = v.BindMemberInfos[:0]
	v.ChannelId = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.H5ListUrl = ""
	v.HotelListUrl = ""
	v.RequestId = ""
	v.TotalResults = 0
	v.Version = 0
	v.Error = false
	v.HasNext = false
	v.Success = false
	poolHotelPriceResultSet.Put(v)
}
