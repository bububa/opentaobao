package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbaseinfogetAPIRequest 酒店基础信息查询接口 API请求
// taobao.xhotel.baseinfo.get
//
// 酒店基础信息(酒店/房型/房价定义)查询接口， 包括 酒店房型可售, 以及 hid 下 的标准房型列表
type TaobaoxhotelbaseinfogetAPIRequest struct {
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

// NewTaobaoxhotelbaseinfogetRequest 初始化TaobaoxhotelbaseinfogetAPIRequest对象
func NewTaobaoxhotelbaseinfogetRequest() *TaobaoxhotelbaseinfogetAPIRequest {
	return &TaobaoxhotelbaseinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbaseinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.baseinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbaseinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbaseinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutHid is OutHid Setter
// 推荐使用卖家系统中的酒店ID。
func (r *TaobaoxhotelbaseinfogetAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoxhotelbaseinfogetAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetVendor is Vendor Setter
// 用于标示该酒店发布的渠道信息
func (r *TaobaoxhotelbaseinfogetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelbaseinfogetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetJsonHotelSellerInvQuery is JsonHotelSellerInvQuery Setter
// 在查询酒店房型可售详情 时的入参JSON , {@link com.taobao.trip.hpc.client.query.HotelSellerInvQuery}
func (r *TaobaoxhotelbaseinfogetAPIRequest) SetJsonHotelSellerInvQuery(_jsonHotelSellerInvQuery string) error {
	r._jsonHotelSellerInvQuery = _jsonHotelSellerInvQuery
	r.Set("json_hotel_seller_inv_query", _jsonHotelSellerInvQuery)
	return nil
}

// GetJsonHotelSellerInvQuery JsonHotelSellerInvQuery Getter
func (r TaobaoxhotelbaseinfogetAPIRequest) GetJsonHotelSellerInvQuery() string {
	return r._jsonHotelSellerInvQuery
}

// SetHid is Hid Setter
// 淘宝酒店ID
func (r *TaobaoxhotelbaseinfogetAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoxhotelbaseinfogetAPIRequest) GetHid() int64 {
	return r._hid
}

// SetIsNeedRatePlan is IsNeedRatePlan Setter
// 是否需要房价基本信息（false为不需要），默认为需要
func (r *TaobaoxhotelbaseinfogetAPIRequest) SetIsNeedRatePlan(_isNeedRatePlan bool) error {
	r._isNeedRatePlan = _isNeedRatePlan
	r.Set("is_need_rate_plan", _isNeedRatePlan)
	return nil
}

// GetIsNeedRatePlan IsNeedRatePlan Getter
func (r TaobaoxhotelbaseinfogetAPIRequest) GetIsNeedRatePlan() bool {
	return r._isNeedRatePlan
}

// SetIsNeedRoomType is IsNeedRoomType Setter
// 是否需要房型基本信息（false为不需要），默认为需要
func (r *TaobaoxhotelbaseinfogetAPIRequest) SetIsNeedRoomType(_isNeedRoomType bool) error {
	r._isNeedRoomType = _isNeedRoomType
	r.Set("is_need_room_type", _isNeedRoomType)
	return nil
}

// GetIsNeedRoomType IsNeedRoomType Getter
func (r TaobaoxhotelbaseinfogetAPIRequest) GetIsNeedRoomType() bool {
	return r._isNeedRoomType
}

// SetNeedSRoomTypeList is NeedSRoomTypeList Setter
// 是否需要 根据 hid 查询 标准房型列表
func (r *TaobaoxhotelbaseinfogetAPIRequest) SetNeedSRoomTypeList(_needSRoomTypeList bool) error {
	r._needSRoomTypeList = _needSRoomTypeList
	r.Set("need_s_room_type_list", _needSRoomTypeList)
	return nil
}

// GetNeedSRoomTypeList NeedSRoomTypeList Getter
func (r TaobaoxhotelbaseinfogetAPIRequest) GetNeedSRoomTypeList() bool {
	return r._needSRoomTypeList
}

// SetNeedHotelDynamicInfo is NeedHotelDynamicInfo Setter
// 是否需要酒店房型可售详情
func (r *TaobaoxhotelbaseinfogetAPIRequest) SetNeedHotelDynamicInfo(_needHotelDynamicInfo bool) error {
	r._needHotelDynamicInfo = _needHotelDynamicInfo
	r.Set("need_hotel_dynamic_info", _needHotelDynamicInfo)
	return nil
}

// GetNeedHotelDynamicInfo NeedHotelDynamicInfo Getter
func (r TaobaoxhotelbaseinfogetAPIRequest) GetNeedHotelDynamicInfo() bool {
	return r._needHotelDynamicInfo
}
