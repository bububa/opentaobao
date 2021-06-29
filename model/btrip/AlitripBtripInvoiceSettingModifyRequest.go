package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票变更 API请求
alitrip.btrip.invoice.setting.modify

发票变更
*/
type AlitripBtripInvoiceSettingModifyRequest struct {
    model.Params
    // 请求对象
    rq   *OpenInvoiceModifyAndNewRq
}

// 初始化AlitripBtripInvoiceSettingModifyRequest对象
func NewAlitripBtripInvoiceSettingModifyRequest() *AlitripBtripInvoiceSettingModifyRequest{
    return &AlitripBtripInvoiceSettingModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSettingModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.setting.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSettingModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripInvoiceSettingModifyRequest) SetRq(rq *OpenInvoiceModifyAndNewRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripInvoiceSettingModifyRequest) GetRq() *OpenInvoiceModifyAndNewRq {
    return r.rq
}
