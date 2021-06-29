package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店基础信息查询接口 API请求
taobao.xhotel.baseinfo.get

酒店基础信息(酒店/房型/房价定义)查询接口， 包括 酒店房型可售, 以及 hid 下 的标准房型列表
*/
type TaobaoXhotelBaseinfoGetRequest struct {
    model.Params
    // 淘宝酒店ID
    _hid   int64
    // 推荐使用卖家系统中的酒店ID。
    _outHid   string
    // 用于标示该酒店发布的渠道信息
    _vendor   string
    // 是否需要房价基本信息（false为不需要），默认为需要
    _isNeedRatePlan   bool
    // 是否需要房型基本信息（false为不需要），默认为需要
    _isNeedRoomType   bool
    // 是否需要 根据 hid 查询 标准房型列表
    _needSRoomTypeList   bool
    // 是否需要酒店房型可售详情
    _needHotelDynamicInfo   bool
    // 在查询酒店房型可售详情 时的入参JSON , {@link com.taobao.trip.hpc.client.query.HotelSellerInvQuery}
    _jsonHotelSellerInvQuery   string
}

// 初始化TaobaoXhotelBaseinfoGetRequest对象
func NewTaobaoXhotelBaseinfoGetRequest() *TaobaoXhotelBaseinfoGetRequest{
    return &TaobaoXhotelBaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBaseinfoGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.baseinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// 淘宝酒店ID
func (r *TaobaoXhotelBaseinfoGetRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelBaseinfoGetRequest) GetHid() int64 {
    return r._hid
}
// OutHid Setter
// 推荐使用卖家系统中的酒店ID。
func (r *TaobaoXhotelBaseinfoGetRequest) SetOutHid(_outHid string) error {
    r._outHid = _outHid
    r.Set("out_hid", _outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelBaseinfoGetRequest) GetOutHid() string {
    return r._outHid
}
// Vendor Setter
// 用于标示该酒店发布的渠道信息
func (r *TaobaoXhotelBaseinfoGetRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelBaseinfoGetRequest) GetVendor() string {
    return r._vendor
}
// IsNeedRatePlan Setter
// 是否需要房价基本信息（false为不需要），默认为需要
func (r *TaobaoXhotelBaseinfoGetRequest) SetIsNeedRatePlan(_isNeedRatePlan bool) error {
    r._isNeedRatePlan = _isNeedRatePlan
    r.Set("is_need_rate_plan", _isNeedRatePlan)
    return nil
}

// IsNeedRatePlan Getter
func (r TaobaoXhotelBaseinfoGetRequest) GetIsNeedRatePlan() bool {
    return r._isNeedRatePlan
}
// IsNeedRoomType Setter
// 是否需要房型基本信息（false为不需要），默认为需要
func (r *TaobaoXhotelBaseinfoGetRequest) SetIsNeedRoomType(_isNeedRoomType bool) error {
    r._isNeedRoomType = _isNeedRoomType
    r.Set("is_need_room_type", _isNeedRoomType)
    return nil
}

// IsNeedRoomType Getter
func (r TaobaoXhotelBaseinfoGetRequest) GetIsNeedRoomType() bool {
    return r._isNeedRoomType
}
// NeedSRoomTypeList Setter
// 是否需要 根据 hid 查询 标准房型列表
func (r *TaobaoXhotelBaseinfoGetRequest) SetNeedSRoomTypeList(_needSRoomTypeList bool) error {
    r._needSRoomTypeList = _needSRoomTypeList
    r.Set("need_s_room_type_list", _needSRoomTypeList)
    return nil
}

// NeedSRoomTypeList Getter
func (r TaobaoXhotelBaseinfoGetRequest) GetNeedSRoomTypeList() bool {
    return r._needSRoomTypeList
}
// NeedHotelDynamicInfo Setter
// 是否需要酒店房型可售详情
func (r *TaobaoXhotelBaseinfoGetRequest) SetNeedHotelDynamicInfo(_needHotelDynamicInfo bool) error {
    r._needHotelDynamicInfo = _needHotelDynamicInfo
    r.Set("need_hotel_dynamic_info", _needHotelDynamicInfo)
    return nil
}

// NeedHotelDynamicInfo Getter
func (r TaobaoXhotelBaseinfoGetRequest) GetNeedHotelDynamicInfo() bool {
    return r._needHotelDynamicInfo
}
// JsonHotelSellerInvQuery Setter
// 在查询酒店房型可售详情 时的入参JSON , {@link com.taobao.trip.hpc.client.query.HotelSellerInvQuery}
func (r *TaobaoXhotelBaseinfoGetRequest) SetJsonHotelSellerInvQuery(_jsonHotelSellerInvQuery string) error {
    r._jsonHotelSellerInvQuery = _jsonHotelSellerInvQuery
    r.Set("json_hotel_seller_inv_query", _jsonHotelSellerInvQuery)
    return nil
}

// JsonHotelSellerInvQuery Getter
func (r TaobaoXhotelBaseinfoGetRequest) GetJsonHotelSellerInvQuery() string {
    return r._jsonHotelSellerInvQuery
}
