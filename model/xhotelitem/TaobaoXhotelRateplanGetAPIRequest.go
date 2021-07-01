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
type TaobaoXhotelRateplanGetAPIRequest struct {
    model.Params
    // 废弃，使用rateplan_code
    _rpid   int64
    // 卖家自己系统的Code，简称RateCode
    _rateplanCode   string
    // 系统商，一般不填写，使用须申请
    _vendor   string
}

// 初始化TaobaoXhotelRateplanGetAPIRequest对象
func NewTaobaoXhotelRateplanGetRequest() *TaobaoXhotelRateplanGetAPIRequest{
    return &TaobaoXhotelRateplanGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateplanGetAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.rateplan.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateplanGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rpid Setter
// 废弃，使用rateplan_code
func (r *TaobaoXhotelRateplanGetAPIRequest) SetRpid(_rpid int64) error {
    r._rpid = _rpid
    r.Set("rpid", _rpid)
    return nil
}

// Rpid Getter
func (r TaobaoXhotelRateplanGetAPIRequest) GetRpid() int64 {
    return r._rpid
}
// RateplanCode Setter
// 卖家自己系统的Code，简称RateCode
func (r *TaobaoXhotelRateplanGetAPIRequest) SetRateplanCode(_rateplanCode string) error {
    r._rateplanCode = _rateplanCode
    r.Set("rateplan_code", _rateplanCode)
    return nil
}

// RateplanCode Getter
func (r TaobaoXhotelRateplanGetAPIRequest) GetRateplanCode() string {
    return r._rateplanCode
}
// Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelRateplanGetAPIRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRateplanGetAPIRequest) GetVendor() string {
    return r._vendor
}
