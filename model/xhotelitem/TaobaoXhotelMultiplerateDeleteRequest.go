package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
复杂价格删除 APIRequest
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

func NewTaobaoXhotelMultiplerateDeleteRequest() *TaobaoXhotelMultiplerateDeleteRequest{
    return &TaobaoXhotelMultiplerateDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelMultiplerateDeleteRequest) GetApiMethodName() string {
    return "taobao.xhotel.multiplerate.delete"
}

func (r TaobaoXhotelMultiplerateDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelMultiplerateDeleteRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelMultiplerateDeleteRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelMultiplerateDeleteRequest) SetRateplanCode(rateplanCode string) error {
    r.rateplanCode = rateplanCode
    r.Set("rateplan_code", rateplanCode)
    return nil
}

func (r TaobaoXhotelMultiplerateDeleteRequest) GetRateplanCode() string {
    return r.rateplanCode
}

func (r *TaobaoXhotelMultiplerateDeleteRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

func (r TaobaoXhotelMultiplerateDeleteRequest) GetOutRid() string {
    return r.outRid
}

func (r *TaobaoXhotelMultiplerateDeleteRequest) SetOccupancy(occupancy int64) error {
    r.occupancy = occupancy
    r.Set("occupancy", occupancy)
    return nil
}

func (r TaobaoXhotelMultiplerateDeleteRequest) GetOccupancy() int64 {
    return r.occupancy
}

func (r *TaobaoXhotelMultiplerateDeleteRequest) SetLengthofstay(lengthofstay int64) error {
    r.lengthofstay = lengthofstay
    r.Set("lengthofstay", lengthofstay)
    return nil
}

func (r TaobaoXhotelMultiplerateDeleteRequest) GetLengthofstay() int64 {
    return r.lengthofstay
}

