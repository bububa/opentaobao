package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票设置 API请求
alitrip.btrip.invoice.setting.add

发票设置
*/
type AlitripBtripInvoiceSettingAddRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenInvoiceModifyAndNewRq
}

// 初始化AlitripBtripInvoiceSettingAddRequest对象
func NewAlitripBtripInvoiceSettingAddRequest() *AlitripBtripInvoiceSettingAddRequest{
    return &AlitripBtripInvoiceSettingAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSettingAddRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.setting.add"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSettingAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripInvoiceSettingAddRequest) SetRq(_rq *OpenInvoiceModifyAndNewRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripInvoiceSettingAddRequest) GetRq() *OpenInvoiceModifyAndNewRq {
    return r._rq
}
