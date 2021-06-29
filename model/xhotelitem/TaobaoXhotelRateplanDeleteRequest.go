package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价格计划rateplan删除 API请求
taobao.xhotel.rateplan.delete

酒店产品库rateplan删除，同时删除级联的rate，请谨慎使用
*/
type TaobaoXhotelRateplanDeleteRequest struct {
    model.Params
    // ratepland标识
    _rpId   int64
    // 系统商，一般不用填写，使用须申请
    _vendor   string
    // 商家价格政策编码
    _rateplanCode   string
}

// 初始化TaobaoXhotelRateplanDeleteRequest对象
func NewTaobaoXhotelRateplanDeleteRequest() *TaobaoXhotelRateplanDeleteRequest{
    return &TaobaoXhotelRateplanDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateplanDeleteRequest) GetApiMethodName() string {
    return "taobao.xhotel.rateplan.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateplanDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RpId Setter
// ratepland标识
func (r *TaobaoXhotelRateplanDeleteRequest) SetRpId(_rpId int64) error {
    r._rpId = _rpId
    r.Set("rp_id", _rpId)
    return nil
}

// RpId Getter
func (r TaobaoXhotelRateplanDeleteRequest) GetRpId() int64 {
    return r._rpId
}
// Vendor Setter
// 系统商，一般不用填写，使用须申请
func (r *TaobaoXhotelRateplanDeleteRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRateplanDeleteRequest) GetVendor() string {
    return r._vendor
}
// RateplanCode Setter
// 商家价格政策编码
func (r *TaobaoXhotelRateplanDeleteRequest) SetRateplanCode(_rateplanCode string) error {
    r._rateplanCode = _rateplanCode
    r.Set("rateplan_code", _rateplanCode)
    return nil
}

// RateplanCode Getter
func (r TaobaoXhotelRateplanDeleteRequest) GetRateplanCode() string {
    return r._rateplanCode
}
