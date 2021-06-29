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
    appId   string
    // 外部订单号
    partnerOrderCode   string
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
func (r *AlibabaEleFengniaoShippingorderEventRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r AlibabaEleFengniaoShippingorderEventRequest) GetAppId() string {
    return r.appId
}
// PartnerOrderCode Setter
// 外部订单号
func (r *AlibabaEleFengniaoShippingorderEventRequest) SetPartnerOrderCode(partnerOrderCode string) error {
    r.partnerOrderCode = partnerOrderCode
    r.Set("partner_order_code", partnerOrderCode)
    return nil
}

// PartnerOrderCode Getter
func (r AlibabaEleFengniaoShippingorderEventRequest) GetPartnerOrderCode() string {
    return r.partnerOrderCode
}
