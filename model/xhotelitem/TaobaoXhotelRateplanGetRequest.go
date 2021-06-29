package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价格计划rateplan查询 APIRequest
taobao.xhotel.rateplan.get

酒店产品库rateplan查询
*/
type TaobaoXhotelRateplanGetRequest struct {
    model.Params

    // 废弃，使用rateplan_code
    rpid   int64 

    // 卖家自己系统的Code，简称RateCode
    rateplanCode   string 

    // 系统商，一般不填写，使用须申请
    vendor   string 

}

func NewTaobaoXhotelRateplanGetRequest() *TaobaoXhotelRateplanGetRequest{
    return &TaobaoXhotelRateplanGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRateplanGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.rateplan.get"
}

func (r TaobaoXhotelRateplanGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRateplanGetRequest) SetRpid(rpid int64) error {
    r.rpid = rpid
    r.Set("rpid", rpid)
    return nil
}

func (r TaobaoXhotelRateplanGetRequest) GetRpid() int64 {
    return r.rpid
}

func (r *TaobaoXhotelRateplanGetRequest) SetRateplanCode(rateplanCode string) error {
    r.rateplanCode = rateplanCode
    r.Set("rateplan_code", rateplanCode)
    return nil
}

func (r TaobaoXhotelRateplanGetRequest) GetRateplanCode() string {
    return r.rateplanCode
}

func (r *TaobaoXhotelRateplanGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelRateplanGetRequest) GetVendor() string {
    return r.vendor
}

