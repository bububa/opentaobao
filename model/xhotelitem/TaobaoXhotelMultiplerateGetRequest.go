package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
复杂房价查询接口 API请求
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

// 初始化TaobaoXhotelMultiplerateGetRequest对象
func NewTaobaoXhotelMultiplerateGetRequest() *TaobaoXhotelMultiplerateGetRequest{
    return &TaobaoXhotelMultiplerateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelMultiplerateGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.multiplerate.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelMultiplerateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nod Setter
// 连住天数(范围1~10)
func (r *TaobaoXhotelMultiplerateGetRequest) SetNod(nod int64) error {
    r.nod = nod
    r.Set("nod", nod)
    return nil
}

// Nod Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetNod() int64 {
    return r.nod
}
// Nop Setter
// 入住人数(范围1~10)
func (r *TaobaoXhotelMultiplerateGetRequest) SetNop(nop int64) error {
    r.nop = nop
    r.Set("nop", nop)
    return nil
}

// Nop Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetNop() int64 {
    return r.nop
}
// RatePlanCode Setter
// 卖家的房价code
func (r *TaobaoXhotelMultiplerateGetRequest) SetRatePlanCode(ratePlanCode string) error {
    r.ratePlanCode = ratePlanCode
    r.Set("rate_plan_code", ratePlanCode)
    return nil
}

// RatePlanCode Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetRatePlanCode() string {
    return r.ratePlanCode
}
// RatePlanId Setter
// 废弃，使用rate_plan_code
func (r *TaobaoXhotelMultiplerateGetRequest) SetRatePlanId(ratePlanId int64) error {
    r.ratePlanId = ratePlanId
    r.Set("rate_plan_id", ratePlanId)
    return nil
}

// RatePlanId Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetRatePlanId() int64 {
    return r.ratePlanId
}
// OutRid Setter
// 卖家的房型code
func (r *TaobaoXhotelMultiplerateGetRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

// OutRid Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetOutRid() string {
    return r.outRid
}
// Gid Setter
// 废弃，使用out_rid
func (r *TaobaoXhotelMultiplerateGetRequest) SetGid(gid int64) error {
    r.gid = gid
    r.Set("gid", gid)
    return nil
}

// Gid Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetGid() int64 {
    return r.gid
}
// Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelMultiplerateGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetVendor() string {
    return r.vendor
}
