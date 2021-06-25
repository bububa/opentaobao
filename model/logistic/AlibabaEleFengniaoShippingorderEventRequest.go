package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询运单事件信息 APIRequest
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

func NewAlibabaEleFengniaoShippingorderEventRequest() *AlibabaEleFengniaoShippingorderEventRequest{
    return &AlibabaEleFengniaoShippingorderEventRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEleFengniaoShippingorderEventRequest) GetApiMethodName() string {
    return "alibaba.ele.fengniao.shippingorder.event"
}

func (r AlibabaEleFengniaoShippingorderEventRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEleFengniaoShippingorderEventRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r AlibabaEleFengniaoShippingorderEventRequest) GetAppId() string {
    return r.appId
}

func (r *AlibabaEleFengniaoShippingorderEventRequest) SetPartnerOrderCode(partnerOrderCode string) error {
    r.partnerOrderCode = partnerOrderCode
    r.Set("partner_order_code", partnerOrderCode)
    return nil
}

func (r AlibabaEleFengniaoShippingorderEventRequest) GetPartnerOrderCode() string {
    return r.partnerOrderCode
}

