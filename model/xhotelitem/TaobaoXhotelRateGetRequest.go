package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店产品库rate查询 API请求
taobao.xhotel.rate.get

酒店产品库rate查询
*/
type TaobaoXhotelRateGetRequest struct {
    model.Params
    // gid酒店商品id
    gid   int64
    // 酒店RPID
    rpid   int64
    // 卖家房型ID, 这是卖家自己系统中的房型ID 注意：需要按照规则组合
    outRid   string
    // 卖家自己系统的Code，简称RateCode
    rateplanCode   string
    // 用于标示该宝贝的售卖渠道信息，允许同一个卖家酒店房型在淘宝系统发布多个售卖渠道的宝贝的价格。
    vendor   string
    // RateID
    rateId   int64
}

// 初始化TaobaoXhotelRateGetRequest对象
func NewTaobaoXhotelRateGetRequest() *TaobaoXhotelRateGetRequest{
    return &TaobaoXhotelRateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.rate.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Gid Setter
// gid酒店商品id
func (r *TaobaoXhotelRateGetRequest) SetGid(gid int64) error {
    r.gid = gid
    r.Set("gid", gid)
    return nil
}

// Gid Getter
func (r TaobaoXhotelRateGetRequest) GetGid() int64 {
    return r.gid
}
// Rpid Setter
// 酒店RPID
func (r *TaobaoXhotelRateGetRequest) SetRpid(rpid int64) error {
    r.rpid = rpid
    r.Set("rpid", rpid)
    return nil
}

// Rpid Getter
func (r TaobaoXhotelRateGetRequest) GetRpid() int64 {
    return r.rpid
}
// OutRid Setter
// 卖家房型ID, 这是卖家自己系统中的房型ID 注意：需要按照规则组合
func (r *TaobaoXhotelRateGetRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

// OutRid Getter
func (r TaobaoXhotelRateGetRequest) GetOutRid() string {
    return r.outRid
}
// RateplanCode Setter
// 卖家自己系统的Code，简称RateCode
func (r *TaobaoXhotelRateGetRequest) SetRateplanCode(rateplanCode string) error {
    r.rateplanCode = rateplanCode
    r.Set("rateplan_code", rateplanCode)
    return nil
}

// RateplanCode Getter
func (r TaobaoXhotelRateGetRequest) GetRateplanCode() string {
    return r.rateplanCode
}
// Vendor Setter
// 用于标示该宝贝的售卖渠道信息，允许同一个卖家酒店房型在淘宝系统发布多个售卖渠道的宝贝的价格。
func (r *TaobaoXhotelRateGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRateGetRequest) GetVendor() string {
    return r.vendor
}
// RateId Setter
// RateID
func (r *TaobaoXhotelRateGetRequest) SetRateId(rateId int64) error {
    r.rateId = rateId
    r.Set("rate_id", rateId)
    return nil
}

// RateId Getter
func (r TaobaoXhotelRateGetRequest) GetRateId() int64 {
    return r.rateId
}
