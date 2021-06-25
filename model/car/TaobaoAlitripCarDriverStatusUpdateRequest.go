package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
司机服务状态更新接口 APIRequest
taobao.alitrip.car.driver.status.update

飞猪用车业务回调接口，用于服务商实时回传更新司机当前服务状态
*/
type TaobaoAlitripCarDriverStatusUpdateRequest struct {
    model.Params

    // 飞猪订单id
    orderId   string 

    // 服务商订单id
    thirdOrderId   string 

    // 服务商标识，由飞猪提供给到各服务商
    providerId   string 

    // 司机服务状态。1-司机已出发，2-司机已到达，3-司机已开始服务
    status   int64 

    // 状态变更相应时间（如司机出发时间、司机到达时间、服务开始时间），格式：yyyy-mm-dd hh:mm:ss
    time   string 

    // 可选，卖家id
    sellerId   string 

    // 0:接送机 1：实时打车 2：租车(不传值默认为0)
    useType   int64 

}

func NewTaobaoAlitripCarDriverStatusUpdateRequest() *TaobaoAlitripCarDriverStatusUpdateRequest{
    return &TaobaoAlitripCarDriverStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.driver.status.update"
}

func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetOrderId() string {
    return r.orderId
}

func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetThirdOrderId(thirdOrderId string) error {
    r.thirdOrderId = thirdOrderId
    r.Set("third_order_id", thirdOrderId)
    return nil
}

func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetThirdOrderId() string {
    return r.thirdOrderId
}

func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetProviderId(providerId string) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetProviderId() string {
    return r.providerId
}

func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetTime(time string) error {
    r.time = time
    r.Set("time", time)
    return nil
}

func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetTime() string {
    return r.time
}

func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetSellerId(sellerId string) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetSellerId() string {
    return r.sellerId
}

func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetUseType(useType int64) error {
    r.useType = useType
    r.Set("use_type", useType)
    return nil
}

func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetUseType() int64 {
    return r.useType
}

