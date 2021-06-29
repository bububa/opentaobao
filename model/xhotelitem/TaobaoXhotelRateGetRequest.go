package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店产品库rate查询 APIRequest
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

func NewTaobaoXhotelRateGetRequest() *TaobaoXhotelRateGetRequest{
    return &TaobaoXhotelRateGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRateGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.rate.get"
}

func (r TaobaoXhotelRateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRateGetRequest) SetGid(gid int64) error {
    r.gid = gid
    r.Set("gid", gid)
    return nil
}

func (r TaobaoXhotelRateGetRequest) GetGid() int64 {
    return r.gid
}

func (r *TaobaoXhotelRateGetRequest) SetRpid(rpid int64) error {
    r.rpid = rpid
    r.Set("rpid", rpid)
    return nil
}

func (r TaobaoXhotelRateGetRequest) GetRpid() int64 {
    return r.rpid
}

func (r *TaobaoXhotelRateGetRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

func (r TaobaoXhotelRateGetRequest) GetOutRid() string {
    return r.outRid
}

func (r *TaobaoXhotelRateGetRequest) SetRateplanCode(rateplanCode string) error {
    r.rateplanCode = rateplanCode
    r.Set("rateplan_code", rateplanCode)
    return nil
}

func (r TaobaoXhotelRateGetRequest) GetRateplanCode() string {
    return r.rateplanCode
}

func (r *TaobaoXhotelRateGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelRateGetRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelRateGetRequest) SetRateId(rateId int64) error {
    r.rateId = rateId
    r.Set("rate_id", rateId)
    return nil
}

func (r TaobaoXhotelRateGetRequest) GetRateId() int64 {
    return r.rateId
}

