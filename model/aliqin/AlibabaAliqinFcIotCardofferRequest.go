package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物联网卡上订购的offer API请求
alibaba.aliqin.fc.iot.cardoffer

查询物联网卡上订购的offer
*/
type AlibabaAliqinFcIotCardofferAPIRequest struct {
    model.Params
    // 具体ICCID的值
    _billreal   string
    // ICCID
    _billsource   string
}

// 初始化AlibabaAliqinFcIotCardofferAPIRequest对象
func NewAlibabaAliqinFcIotCardofferRequest() *AlibabaAliqinFcIotCardofferAPIRequest{
    return &AlibabaAliqinFcIotCardofferAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotCardofferAPIRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.cardoffer"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotCardofferAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Billreal Setter
// 具体ICCID的值
func (r *AlibabaAliqinFcIotCardofferAPIRequest) SetBillreal(_billreal string) error {
    r._billreal = _billreal
    r.Set("billreal", _billreal)
    return nil
}

// Billreal Getter
func (r AlibabaAliqinFcIotCardofferAPIRequest) GetBillreal() string {
    return r._billreal
}
// Billsource Setter
// ICCID
func (r *AlibabaAliqinFcIotCardofferAPIRequest) SetBillsource(_billsource string) error {
    r._billsource = _billsource
    r.Set("billsource", _billsource)
    return nil
}

// Billsource Getter
func (r AlibabaAliqinFcIotCardofferAPIRequest) GetBillsource() string {
    return r._billsource
}
