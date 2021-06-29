package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
复杂价格删除 API请求
taobao.xhotel.multiplerate.delete

酒店产品库rate删除
*/
type TaobaoXhotelMultiplerateDeleteRequest struct {
    model.Params
    // 渠道，和推送房价所使用的渠道保持一致
    vendor   string
    // 商家价格政策编码
    rateplanCode   string
    // 商家房型编码
    outRid   string
    // 连住天数
    occupancy   int64
    // 入住人数
    lengthofstay   int64
}

// 初始化TaobaoXhotelMultiplerateDeleteRequest对象
func NewTaobaoXhotelMultiplerateDeleteRequest() *TaobaoXhotelMultiplerateDeleteRequest{
    return &TaobaoXhotelMultiplerateDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelMultiplerateDeleteRequest) GetApiMethodName() string {
    return "taobao.xhotel.multiplerate.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelMultiplerateDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Vendor Setter
// 渠道，和推送房价所使用的渠道保持一致
func (r *TaobaoXhotelMultiplerateDeleteRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelMultiplerateDeleteRequest) GetVendor() string {
    return r.vendor
}
// RateplanCode Setter
// 商家价格政策编码
func (r *TaobaoXhotelMultiplerateDeleteRequest) SetRateplanCode(rateplanCode string) error {
    r.rateplanCode = rateplanCode
    r.Set("rateplan_code", rateplanCode)
    return nil
}

// RateplanCode Getter
func (r TaobaoXhotelMultiplerateDeleteRequest) GetRateplanCode() string {
    return r.rateplanCode
}
// OutRid Setter
// 商家房型编码
func (r *TaobaoXhotelMultiplerateDeleteRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

// OutRid Getter
func (r TaobaoXhotelMultiplerateDeleteRequest) GetOutRid() string {
    return r.outRid
}
// Occupancy Setter
// 连住天数
func (r *TaobaoXhotelMultiplerateDeleteRequest) SetOccupancy(occupancy int64) error {
    r.occupancy = occupancy
    r.Set("occupancy", occupancy)
    return nil
}

// Occupancy Getter
func (r TaobaoXhotelMultiplerateDeleteRequest) GetOccupancy() int64 {
    return r.occupancy
}
// Lengthofstay Setter
// 入住人数
func (r *TaobaoXhotelMultiplerateDeleteRequest) SetLengthofstay(lengthofstay int64) error {
    r.lengthofstay = lengthofstay
    r.Set("lengthofstay", lengthofstay)
    return nil
}

// Lengthofstay Getter
func (r TaobaoXhotelMultiplerateDeleteRequest) GetLengthofstay() int64 {
    return r.lengthofstay
}
