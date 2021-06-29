package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
复杂房价查询接口 APIRequest
taobao.xhotel.multiplerate.get

查询复杂房价，支持通过入住人数，连住天数，商品信息，房价信息查询
*/
type TaobaoXhotelMultiplerateGetRequest struct {
    model.Params

    // 连住天数(范围1~10)
    nod   int64 

    // 入住人数(范围1~10)
    nop   int64 

    // 卖家的房价code
    ratePlanCode   string 

    // 废弃，使用rate_plan_code
    ratePlanId   int64 

    // 卖家的房型code
    outRid   string 

    // 废弃，使用out_rid
    gid   int64 

    // 系统商，一般不填写，使用须申请
    vendor   string 

}

func NewTaobaoXhotelMultiplerateGetRequest() *TaobaoXhotelMultiplerateGetRequest{
    return &TaobaoXhotelMultiplerateGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelMultiplerateGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.multiplerate.get"
}

func (r TaobaoXhotelMultiplerateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelMultiplerateGetRequest) SetNod(nod int64) error {
    r.nod = nod
    r.Set("nod", nod)
    return nil
}

func (r TaobaoXhotelMultiplerateGetRequest) GetNod() int64 {
    return r.nod
}

func (r *TaobaoXhotelMultiplerateGetRequest) SetNop(nop int64) error {
    r.nop = nop
    r.Set("nop", nop)
    return nil
}

func (r TaobaoXhotelMultiplerateGetRequest) GetNop() int64 {
    return r.nop
}

func (r *TaobaoXhotelMultiplerateGetRequest) SetRatePlanCode(ratePlanCode string) error {
    r.ratePlanCode = ratePlanCode
    r.Set("rate_plan_code", ratePlanCode)
    return nil
}

func (r TaobaoXhotelMultiplerateGetRequest) GetRatePlanCode() string {
    return r.ratePlanCode
}

func (r *TaobaoXhotelMultiplerateGetRequest) SetRatePlanId(ratePlanId int64) error {
    r.ratePlanId = ratePlanId
    r.Set("rate_plan_id", ratePlanId)
    return nil
}

func (r TaobaoXhotelMultiplerateGetRequest) GetRatePlanId() int64 {
    return r.ratePlanId
}

func (r *TaobaoXhotelMultiplerateGetRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

func (r TaobaoXhotelMultiplerateGetRequest) GetOutRid() string {
    return r.outRid
}

func (r *TaobaoXhotelMultiplerateGetRequest) SetGid(gid int64) error {
    r.gid = gid
    r.Set("gid", gid)
    return nil
}

func (r TaobaoXhotelMultiplerateGetRequest) GetGid() int64 {
    return r.gid
}

func (r *TaobaoXhotelMultiplerateGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelMultiplerateGetRequest) GetVendor() string {
    return r.vendor
}

