package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流信息回传接口 API请求
alibaba.alihealth.nr.delivery.history.save

商家ERP回传物流信息
*/
type AlibabaAlihealthNrDeliveryHistorySaveRequest struct {
    model.Params
    // 入参
    _deliveryEvent   *DeliveryEventDTO
}

// 初始化AlibabaAlihealthNrDeliveryHistorySaveRequest对象
func NewAlibabaAlihealthNrDeliveryHistorySaveRequest() *AlibabaAlihealthNrDeliveryHistorySaveRequest{
    return &AlibabaAlihealthNrDeliveryHistorySaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrDeliveryHistorySaveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.delivery.history.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrDeliveryHistorySaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeliveryEvent Setter
// 入参
func (r *AlibabaAlihealthNrDeliveryHistorySaveRequest) SetDeliveryEvent(_deliveryEvent *DeliveryEventDTO) error {
    r._deliveryEvent = _deliveryEvent
    r.Set("delivery_event", _deliveryEvent)
    return nil
}

// DeliveryEvent Getter
func (r AlibabaAlihealthNrDeliveryHistorySaveRequest) GetDeliveryEvent() *DeliveryEventDTO {
    return r._deliveryEvent
}
