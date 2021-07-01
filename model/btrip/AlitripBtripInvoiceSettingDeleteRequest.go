package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票删除 API请求
alitrip.btrip.invoice.setting.delete

发票删除
*/
type AlitripBtripInvoiceSettingDeleteAPIRequest struct {
    model.Params
    // 入参
    _rq   *OpenInvoiceDeleteRq
}

// 初始化AlitripBtripInvoiceSettingDeleteAPIRequest对象
func NewAlitripBtripInvoiceSettingDeleteRequest() *AlitripBtripInvoiceSettingDeleteAPIRequest{
    return &AlitripBtripInvoiceSettingDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSettingDeleteAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.setting.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSettingDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参
func (r *AlitripBtripInvoiceSettingDeleteAPIRequest) SetRq(_rq *OpenInvoiceDeleteRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripInvoiceSettingDeleteAPIRequest) GetRq() *OpenInvoiceDeleteRq {
    return r._rq
}
