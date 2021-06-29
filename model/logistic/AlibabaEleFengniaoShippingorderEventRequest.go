package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询运单事件信息 API请求
alibaba.ele.fengniao.shippingorder.event

查询运单事件信息
*/
type AlibabaEleFengniaoShippingorderEventRequest struct {
    model.Params
    // appid
    _appId   string
    // 外部订单号
    _partnerOrderCode   string
}

// 初始化AlibabaEleFengniaoShippingorderEventRequest对象
func NewAlibabaEleFengniaoShippingorderEventRequest() *AlibabaEleFengniaoShippingorderEventRequest{
    return &AlibabaEleFengniaoShippingorderEventRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoShippingorderEventRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.shippingorder.event"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoShippingorderEventRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppId Setter
// appid
func (r *AlibabaEleFengniaoShippingorderEventRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r AlibabaEleFengniaoShippingorderEventRequest) GetAppId() string {
    return r._appId
}
// PartnerOrderCode Setter
// 外部订单号
func (r *AlibabaEleFengniaoShippingorderEventRequest) SetPartnerOrderCode(_partnerOrderCode string) error {
    r._partnerOrderCode = _partnerOrderCode
    r.Set("partner_order_code", _partnerOrderCode)
    return nil
}

// PartnerOrderCode Getter
func (r AlibabaEleFengniaoShippingorderEventRequest) GetPartnerOrderCode() string {
    return r._partnerOrderCode
}
