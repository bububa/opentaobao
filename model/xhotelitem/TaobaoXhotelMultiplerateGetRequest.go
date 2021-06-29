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
    _nod   int64
    // 入住人数(范围1~10)
    _nop   int64
    // 卖家的房价code
    _ratePlanCode   string
    // 废弃，使用rate_plan_code
    _ratePlanId   int64
    // 卖家的房型code
    _outRid   string
    // 废弃，使用out_rid
    _gid   int64
    // 系统商，一般不填写，使用须申请
    _vendor   string
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
func (r *TaobaoXhotelMultiplerateGetRequest) SetNod(_nod int64) error {
    r._nod = _nod
    r.Set("nod", _nod)
    return nil
}

// Nod Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetNod() int64 {
    return r._nod
}
// Nop Setter
// 入住人数(范围1~10)
func (r *TaobaoXhotelMultiplerateGetRequest) SetNop(_nop int64) error {
    r._nop = _nop
    r.Set("nop", _nop)
    return nil
}

// Nop Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetNop() int64 {
    return r._nop
}
// RatePlanCode Setter
// 卖家的房价code
func (r *TaobaoXhotelMultiplerateGetRequest) SetRatePlanCode(_ratePlanCode string) error {
    r._ratePlanCode = _ratePlanCode
    r.Set("rate_plan_code", _ratePlanCode)
    return nil
}

// RatePlanCode Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetRatePlanCode() string {
    return r._ratePlanCode
}
// RatePlanId Setter
// 废弃，使用rate_plan_code
func (r *TaobaoXhotelMultiplerateGetRequest) SetRatePlanId(_ratePlanId int64) error {
    r._ratePlanId = _ratePlanId
    r.Set("rate_plan_id", _ratePlanId)
    return nil
}

// RatePlanId Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetRatePlanId() int64 {
    return r._ratePlanId
}
// OutRid Setter
// 卖家的房型code
func (r *TaobaoXhotelMultiplerateGetRequest) SetOutRid(_outRid string) error {
    r._outRid = _outRid
    r.Set("out_rid", _outRid)
    return nil
}

// OutRid Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetOutRid() string {
    return r._outRid
}
// Gid Setter
// 废弃，使用out_rid
func (r *TaobaoXhotelMultiplerateGetRequest) SetGid(_gid int64) error {
    r._gid = _gid
    r.Set("gid", _gid)
    return nil
}

// Gid Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetGid() int64 {
    return r._gid
}
// Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelMultiplerateGetRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelMultiplerateGetRequest) GetVendor() string {
    return r._vendor
}