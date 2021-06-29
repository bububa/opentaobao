package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价格计划rateplan查询 API请求
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

// 初始化TaobaoXhotelRateplanGetRequest对象
func NewTaobaoXhotelRateplanGetRequest() *TaobaoXhotelRateplanGetRequest{
    return &TaobaoXhotelRateplanGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateplanGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.rateplan.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateplanGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rpid Setter
// 废弃，使用rateplan_code
func (r *TaobaoXhotelRateplanGetRequest) SetRpid(rpid int64) error {
    r.rpid = rpid
    r.Set("rpid", rpid)
    return nil
}

// Rpid Getter
func (r TaobaoXhotelRateplanGetRequest) GetRpid() int64 {
    return r.rpid
}
// RateplanCode Setter
// 卖家自己系统的Code，简称RateCode
func (r *TaobaoXhotelRateplanGetRequest) SetRateplanCode(rateplanCode string) error {
    r.rateplanCode = rateplanCode
    r.Set("rateplan_code", rateplanCode)
    return nil
}

// RateplanCode Getter
func (r TaobaoXhotelRateplanGetRequest) GetRateplanCode() string {
    return r.rateplanCode
}
// Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelRateplanGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRateplanGetRequest) GetVendor() string {
    return r.vendor
}
