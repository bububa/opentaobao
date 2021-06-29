package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
航旅国内租车订单状态更新 API请求
taobao.alitrip.domestic.rent.car.status.update

航旅国内租车订单状态更新
*/
type TaobaoAlitripDomesticRentCarStatusUpdateRequest struct {
    model.Params
    // 121-用车中（用户取车成功） 122-待结算（用户还车成功）
    _status   int64
    // 服务商平台订单号
    _thirdOrderId   string
    // 飞猪平台订单号
    _orderId   string
    // 服务商标识，由飞猪提供给到各服务商
    _providerId   string
}

// 初始化TaobaoAlitripDomesticRentCarStatusUpdateRequest对象
func NewTaobaoAlitripDomesticRentCarStatusUpdateRequest() *TaobaoAlitripDomesticRentCarStatusUpdateRequest{
    return &TaobaoAlitripDomesticRentCarStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetApiMethodName() string {
    return "taobao.alitrip.domestic.rent.car.status.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// 121-用车中（用户取车成功） 122-待结算（用户还车成功）
func (r *TaobaoAlitripDomesticRentCarStatusUpdateRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetStatus() int64 {
    return r._status
}
// ThirdOrderId Setter
// 服务商平台订单号
func (r *TaobaoAlitripDomesticRentCarStatusUpdateRequest) SetThirdOrderId(_thirdOrderId string) error {
    r._thirdOrderId = _thirdOrderId
    r.Set("third_order_id", _thirdOrderId)
    return nil
}

// ThirdOrderId Getter
func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetThirdOrderId() string {
    return r._thirdOrderId
}
// OrderId Setter
// 飞猪平台订单号
func (r *TaobaoAlitripDomesticRentCarStatusUpdateRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetOrderId() string {
    return r._orderId
}
// ProviderId Setter
// 服务商标识，由飞猪提供给到各服务商
func (r *TaobaoAlitripDomesticRentCarStatusUpdateRequest) SetProviderId(_providerId string) error {
    r._providerId = _providerId
    r.Set("provider_id", _providerId)
    return nil
}

// ProviderId Getter
func (r TaobaoAlitripDomesticRentCarStatusUpdateRequest) GetProviderId() string {
    return r._providerId
}
