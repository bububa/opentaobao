package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流信息回传接口 APIRequest
alibaba.alihealth.nr.delivery.history.save

商家ERP回传物流信息
*/
type AlibabaAlihealthNrDeliveryHistorySaveRequest struct {
    model.Params

    // 入参
    deliveryEvent   *DeliveryEventDto 

}

func NewAlibabaAlihealthNrDeliveryHistorySaveRequest() *AlibabaAlihealthNrDeliveryHistorySaveRequest{
    return &AlibabaAlihealthNrDeliveryHistorySaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthNrDeliveryHistorySaveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.delivery.history.save"
}

func (r AlibabaAlihealthNrDeliveryHistorySaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthNrDeliveryHistorySaveRequest) SetDeliveryEvent(deliveryEvent *DeliveryEventDto) error {
    r.deliveryEvent = deliveryEvent
    r.Set("delivery_event", deliveryEvent)
    return nil
}

func (r AlibabaAlihealthNrDeliveryHistorySaveRequest) GetDeliveryEvent() *DeliveryEventDto {
    return r.deliveryEvent
}

