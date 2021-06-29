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
    _vendor   string
    // 商家价格政策编码
    _rateplanCode   string
    // 商家房型编码
    _outRid   string
    // 连住天数
    _occupancy   int64
    // 入住人数
    _lengthofstay   int64
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
func (r *TaobaoXhotelMultiplerateDeleteRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelMultiplerateDeleteRequest) GetVendor() string {
    return r._vendor
}
// RateplanCode Setter
// 商家价格政策编码
func (r *TaobaoXhotelMultiplerateDeleteRequest) SetRateplanCode(_rateplanCode string) error {
    r._rateplanCode = _rateplanCode
    r.Set("rateplan_code", _rateplanCode)
    return nil
}

// RateplanCode Getter
func (r TaobaoXhotelMultiplerateDeleteRequest) GetRateplanCode() string {
    return r._rateplanCode
}
// OutRid Setter
// 商家房型编码
func (r *TaobaoXhotelMultiplerateDeleteRequest) SetOutRid(_outRid string) error {
    r._outRid = _outRid
    r.Set("out_rid", _outRid)
    return nil
}

// OutRid Getter
func (r TaobaoXhotelMultiplerateDeleteRequest) GetOutRid() string {
    return r._outRid
}
// Occupancy Setter
// 连住天数
func (r *TaobaoXhotelMultiplerateDeleteRequest) SetOccupancy(_occupancy int64) error {
    r._occupancy = _occupancy
    r.Set("occupancy", _occupancy)
    return nil
}

// Occupancy Getter
func (r TaobaoXhotelMultiplerateDeleteRequest) GetOccupancy() int64 {
    return r._occupancy
}
// Lengthofstay Setter
// 入住人数
func (r *TaobaoXhotelMultiplerateDeleteRequest) SetLengthofstay(_lengthofstay int64) error {
    r._lengthofstay = _lengthofstay
    r.Set("lengthofstay", _lengthofstay)
    return nil
}

// Lengthofstay Getter
func (r TaobaoXhotelMultiplerateDeleteRequest) GetLengthofstay() int64 {
    return r._lengthofstay
}
