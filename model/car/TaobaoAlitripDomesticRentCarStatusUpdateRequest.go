package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
航旅国内租车订单状态更新 APIRequest
taobao.alitrip.domestic.rent.car.status.update

航旅国内租车订单状态更新
*/
type TaobaoAlitripDomesticRentCarStatusUpdateRequest struct {
    model.Params

    // 121-用车中（用户取车成功） 122-待结算（用户还车成功）
    status   int64 

    // 服务商平台订单号
    thirdOrderId   string 

    // 飞猪平台订单号
    orderId   string 

    // 服务商标识，由飞猪提供给到各服务商
    providerId   string 

}

func NewTaobaoAlitripDomesticRentCarStatusUpdateRequest() *TaobaoAlitripDomesticRentCarStatusUpdateRequest{
    return &TaobaoAlitripDomesticRentCarStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetApiMethodName() string {
    return "taobao.alitrip.domestic.rent.car.status.update"
}

func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripDomesticRentCarStatusUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoAlitripDomesticRentCarStatusUpdateRequest) SetThirdOrderId(thirdOrderId string) error {
    r.thirdOrderId = thirdOrderId
    r.Set("third_order_id", thirdOrderId)
    return nil
}

func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetThirdOrderId() string {
    return r.thirdOrderId
}

func (r *TaobaoAlitripDomesticRentCarStatusUpdateRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetOrderId() string {
    return r.orderId
}

func (r *TaobaoAlitripDomesticRentCarStatusUpdateRequest) SetProviderId(providerId string) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetProviderId() string {
    return r.providerId
}

