package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店基础信息查询接口 APIRequest
taobao.xhotel.baseinfo.get

酒店基础信息(酒店/房型/房价定义)查询接口， 包括 酒店房型可售, 以及 hid 下 的标准房型列表
*/
type TaobaoXhotelBaseinfoGetRequest struct {
    model.Params

    // 淘宝酒店ID
    hid   int64 

    // 推荐使用卖家系统中的酒店ID。
    outHid   string 

    // 用于标示该酒店发布的渠道信息
    vendor   string 

    // 是否需要房价基本信息（false为不需要），默认为需要
    isNeedRatePlan   bool 

    // 是否需要房型基本信息（false为不需要），默认为需要
    isNeedRoomType   bool 

    // 是否需要 根据 hid 查询 标准房型列表
    needSRoomTypeList   bool 

    // 是否需要酒店房型可售详情
    needHotelDynamicInfo   bool 

    // 在查询酒店房型可售详情 时的入参JSON , {@link com.taobao.trip.hpc.client.query.HotelSellerInvQuery}
    jsonHotelSellerInvQuery   string 

}

func NewTaobaoXhotelBaseinfoGetRequest() *TaobaoXhotelBaseinfoGetRequest{
    return &TaobaoXhotelBaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelBaseinfoGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.baseinfo.get"
}

func (r TaobaoXhotelBaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelBaseinfoGetRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

func (r TaobaoXhotelBaseinfoGetRequest) GetHid() int64 {
    return r.hid
}

func (r *TaobaoXhotelBaseinfoGetRequest) SetOutHid(outHid string) error {
    r.outHid = outHid
    r.Set("out_hid", outHid)
    return nil
}

func (r TaobaoXhotelBaseinfoGetRequest) GetOutHid() string {
    return r.outHid
}

func (r *TaobaoXhotelBaseinfoGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelBaseinfoGetRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelBaseinfoGetRequest) SetIsNeedRatePlan(isNeedRatePlan bool) error {
    r.isNeedRatePlan = isNeedRatePlan
    r.Set("is_need_rate_plan", isNeedRatePlan)
    return nil
}

func (r TaobaoXhotelBaseinfoGetRequest) GetIsNeedRatePlan() bool {
    return r.isNeedRatePlan
}

func (r *TaobaoXhotelBaseinfoGetRequest) SetIsNeedRoomType(isNeedRoomType bool) error {
    r.isNeedRoomType = isNeedRoomType
    r.Set("is_need_room_type", isNeedRoomType)
    return nil
}

func (r TaobaoXhotelBaseinfoGetRequest) GetIsNeedRoomType() bool {
    return r.isNeedRoomType
}

func (r *TaobaoXhotelBaseinfoGetRequest) SetNeedSRoomTypeList(needSRoomTypeList bool) error {
    r.needSRoomTypeList = needSRoomTypeList
    r.Set("need_s_room_type_list", needSRoomTypeList)
    return nil
}

func (r TaobaoXhotelBaseinfoGetRequest) GetNeedSRoomTypeList() bool {
    return r.needSRoomTypeList
}

func (r *TaobaoXhotelBaseinfoGetRequest) SetNeedHotelDynamicInfo(needHotelDynamicInfo bool) error {
    r.needHotelDynamicInfo = needHotelDynamicInfo
    r.Set("need_hotel_dynamic_info", needHotelDynamicInfo)
    return nil
}

func (r TaobaoXhotelBaseinfoGetRequest) GetNeedHotelDynamicInfo() bool {
    return r.needHotelDynamicInfo
}

func (r *TaobaoXhotelBaseinfoGetRequest) SetJsonHotelSellerInvQuery(jsonHotelSellerInvQuery string) error {
    r.jsonHotelSellerInvQuery = jsonHotelSellerInvQuery
    r.Set("json_hotel_seller_inv_query", jsonHotelSellerInvQuery)
    return nil
}

func (r TaobaoXhotelBaseinfoGetRequest) GetJsonHotelSellerInvQuery() string {
    return r.jsonHotelSellerInvQuery
}

