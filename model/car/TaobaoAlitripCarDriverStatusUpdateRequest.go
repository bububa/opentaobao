package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
司机服务状态更新接口 API请求
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

// 初始化TaobaoAlitripCarDriverStatusUpdateRequest对象
func NewTaobaoAlitripCarDriverStatusUpdateRequest() *TaobaoAlitripCarDriverStatusUpdateRequest{
    return &TaobaoAlitripCarDriverStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.driver.status.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 飞猪订单id
func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetOrderId() string {
    return r.orderId
}
// ThirdOrderId Setter
// 服务商订单id
func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetThirdOrderId(thirdOrderId string) error {
    r.thirdOrderId = thirdOrderId
    r.Set("third_order_id", thirdOrderId)
    return nil
}

// ThirdOrderId Getter
func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetThirdOrderId() string {
    return r.thirdOrderId
}
// ProviderId Setter
// 服务商标识，由飞猪提供给到各服务商
func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetProviderId(providerId string) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

// ProviderId Getter
func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetProviderId() string {
    return r.providerId
}
// Status Setter
// 司机服务状态。1-司机已出发，2-司机已到达，3-司机已开始服务
func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetStatus() int64 {
    return r.status
}
// Time Setter
// 状态变更相应时间（如司机出发时间、司机到达时间、服务开始时间），格式：yyyy-mm-dd hh:mm:ss
func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetTime(time string) error {
    r.time = time
    r.Set("time", time)
    return nil
}

// Time Getter
func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetTime() string {
    return r.time
}
// SellerId Setter
// 可选，卖家id
func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetSellerId(sellerId string) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetSellerId() string {
    return r.sellerId
}
// UseType Setter
// 0:接送机 1：实时打车 2：租车(不传值默认为0)
func (r *TaobaoAlitripCarDriverStatusUpdateRequest) SetUseType(useType int64) error {
    r.useType = useType
    r.Set("use_type", useType)
    return nil
}

// UseType Getter
func (r TaobaoAlitripCarDriverStatusUpdateRequest) GetUseType() int64 {
    return r.useType
}
