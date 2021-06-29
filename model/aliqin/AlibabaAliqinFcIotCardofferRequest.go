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
type AlibabaAliqinFcIotCardofferRequest struct {
    model.Params
    // 具体ICCID的值
    _billreal   string
    // ICCID
    _billsource   string
}

// 初始化AlibabaAliqinFcIotCardofferRequest对象
func NewAlibabaAliqinFcIotCardofferRequest() *AlibabaAliqinFcIotCardofferRequest{
    return &AlibabaAliqinFcIotCardofferRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotCardofferRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.cardoffer"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotCardofferRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Billreal Setter
// 具体ICCID的值
func (r *AlibabaAliqinFcIotCardofferRequest) SetBillreal(_billreal string) error {
    r._billreal = _billreal
    r.Set("billreal", _billreal)
    return nil
}

// Billreal Getter
func (r AlibabaAliqinFcIotCardofferRequest) GetBillreal() string {
    return r._billreal
}
// Billsource Setter
// ICCID
func (r *AlibabaAliqinFcIotCardofferRequest) SetBillsource(_billsource string) error {
    r._billsource = _billsource
    r.Set("billsource", _billsource)
    return nil
}

// Billsource Getter
func (r AlibabaAliqinFcIotCardofferRequest) GetBillsource() string {
    return r._billsource
}
