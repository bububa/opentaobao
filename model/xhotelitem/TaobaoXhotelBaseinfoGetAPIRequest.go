package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBaseinfoGetAPIRequest 酒店基础信息查询接口 API请求
// taobao.xhotel.baseinfo.get
//
// 酒店基础信息(酒店/房型/房价定义)查询接口， 包括 酒店房型可售, 以及 hid 下 的标准房型列表
type TaobaoXhotelBaseinfoGetAPIRequest struct {
	model.Params
	// 推荐使用卖家系统中的酒店ID。
	_outHid string
	// 用于标示该酒店发布的渠道信息
	_vendor string
	// 在查询酒店房型可售详情 时的入参JSON , {@link com.taobao.trip.hpc.client.query.HotelSellerInvQuery}
	_jsonHotelSellerInvQuery string
	// 淘宝酒店ID
	_hid int64
	// 是否需要房价基本信息（false为不需要），默认为需要
	_isNeedRatePlan bool
	// 是否需要房型基本信息（false为不需要），默认为需要
	_isNeedRoomType bool
	// 是否需要 根据 hid 查询 标准房型列表
	_needSRoomTypeList bool
	// 是否需要酒店房型可售详情
	_needHotelDynamicInfo bool
}

// NewTaobaoXhotelBaseinfoGetRequest 初始化TaobaoXhotelBaseinfoGetAPIRequest对象
func NewTaobaoXhotelBaseinfoGetRequest() *TaobaoXhotelBaseinfoGetAPIRequest {
	return &TaobaoXhotelBaseinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBaseinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.baseinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBaseinfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOutHid is OutHid Setter
// 推荐使用卖家系统中的酒店ID。
func (r *TaobaoXhotelBaseinfoGetAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoXhotelBaseinfoGetAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetVendor is Vendor Setter
// 用于标示该酒店发布的渠道信息
func (r *TaobaoXhotelBaseinfoGetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelBaseinfoGetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetJsonHotelSellerInvQuery is JsonHotelSellerInvQuery Setter
// 在查询酒店房型可售详情 时的入参JSON , {@link com.taobao.trip.hpc.client.query.HotelSellerInvQuery}
func (r *TaobaoXhotelBaseinfoGetAPIRequest) SetJsonHotelSellerInvQuery(_jsonHotelSellerInvQuery string) error {
	r._jsonHotelSellerInvQuery = _jsonHotelSellerInvQuery
	r.Set("json_hotel_seller_inv_query", _jsonHotelSellerInvQuery)
	return nil
}

// GetJsonHotelSellerInvQuery JsonHotelSellerInvQuery Getter
func (r TaobaoXhotelBaseinfoGetAPIRequest) GetJsonHotelSellerInvQuery() string {
	return r._jsonHotelSellerInvQuery
}

// SetHid is Hid Setter
// 淘宝酒店ID
func (r *TaobaoXhotelBaseinfoGetAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelBaseinfoGetAPIRequest) GetHid() int64 {
	return r._hid
}

// SetIsNeedRatePlan is IsNeedRatePlan Setter
// 是否需要房价基本信息（false为不需要），默认为需要
func (r *TaobaoXhotelBaseinfoGetAPIRequest) SetIsNeedRatePlan(_isNeedRatePlan bool) error {
	r._isNeedRatePlan = _isNeedRatePlan
	r.Set("is_need_rate_plan", _isNeedRatePlan)
	return nil
}

// GetIsNeedRatePlan IsNeedRatePlan Getter
func (r TaobaoXhotelBaseinfoGetAPIRequest) GetIsNeedRatePlan() bool {
	return r._isNeedRatePlan
}

// SetIsNeedRoomType is IsNeedRoomType Setter
// 是否需要房型基本信息（false为不需要），默认为需要
func (r *TaobaoXhotelBaseinfoGetAPIRequest) SetIsNeedRoomType(_isNeedRoomType bool) error {
	r._isNeedRoomType = _isNeedRoomType
	r.Set("is_need_room_type", _isNeedRoomType)
	return nil
}

// GetIsNeedRoomType IsNeedRoomType Getter
func (r TaobaoXhotelBaseinfoGetAPIRequest) GetIsNeedRoomType() bool {
	return r._isNeedRoomType
}

// SetNeedSRoomTypeList is NeedSRoomTypeList Setter
// 是否需要 根据 hid 查询 标准房型列表
func (r *TaobaoXhotelBaseinfoGetAPIRequest) SetNeedSRoomTypeList(_needSRoomTypeList bool) error {
	r._needSRoomTypeList = _needSRoomTypeList
	r.Set("need_s_room_type_list", _needSRoomTypeList)
	return nil
}

// GetNeedSRoomTypeList NeedSRoomTypeList Getter
func (r TaobaoXhotelBaseinfoGetAPIRequest) GetNeedSRoomTypeList() bool {
	return r._needSRoomTypeList
}

// SetNeedHotelDynamicInfo is NeedHotelDynamicInfo Setter
// 是否需要酒店房型可售详情
func (r *TaobaoXhotelBaseinfoGetAPIRequest) SetNeedHotelDynamicInfo(_needHotelDynamicInfo bool) error {
	r._needHotelDynamicInfo = _needHotelDynamicInfo
	r.Set("need_hotel_dynamic_info", _needHotelDynamicInfo)
	return nil
}

// GetNeedHotelDynamicInfo NeedHotelDynamicInfo Getter
func (r TaobaoXhotelBaseinfoGetAPIRequest) GetNeedHotelDynamicInfo() bool {
	return r._needHotelDynamicInfo
}
