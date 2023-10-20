package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripdomesticrentcarstatusupdateAPIRequest 航旅国内租车订单状态更新 API请求
// taobao.alitrip.domestic.rent.car.status.update
//
// 航旅国内租车订单状态更新
type TaobaoalitripdomesticrentcarstatusupdateAPIRequest struct {
	model.Params
	// 服务商平台订单号
	_thirdOrderId string
	// 飞猪平台订单号
	_orderId string
	// 服务商标识，由飞猪提供给到各服务商
	_providerId string
	// 车牌号
	_carNumber string
	// JSON扩展值
	_extra string
	// 121-用车中（用户取车成功） 122-待结算（用户还车成功）
	_status int64
}

// NewTaobaoalitripdomesticrentcarstatusupdateRequest 初始化TaobaoalitripdomesticrentcarstatusupdateAPIRequest对象
func NewTaobaoalitripdomesticrentcarstatusupdateRequest() *TaobaoalitripdomesticrentcarstatusupdateAPIRequest {
	return &TaobaoalitripdomesticrentcarstatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripdomesticrentcarstatusupdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.domestic.rent.car.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripdomesticrentcarstatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripdomesticrentcarstatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetThirdOrderId is ThirdOrderId Setter
// 服务商平台订单号
func (r *TaobaoalitripdomesticrentcarstatusupdateAPIRequest) SetThirdOrderId(_thirdOrderId string) error {
	r._thirdOrderId = _thirdOrderId
	r.Set("third_order_id", _thirdOrderId)
	return nil
}

// GetThirdOrderId ThirdOrderId Getter
func (r TaobaoalitripdomesticrentcarstatusupdateAPIRequest) GetThirdOrderId() string {
	return r._thirdOrderId
}

// SetOrderId is OrderId Setter
// 飞猪平台订单号
func (r *TaobaoalitripdomesticrentcarstatusupdateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoalitripdomesticrentcarstatusupdateAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetProviderId is ProviderId Setter
// 服务商标识，由飞猪提供给到各服务商
func (r *TaobaoalitripdomesticrentcarstatusupdateAPIRequest) SetProviderId(_providerId string) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r TaobaoalitripdomesticrentcarstatusupdateAPIRequest) GetProviderId() string {
	return r._providerId
}

// SetCarNumber is CarNumber Setter
// 车牌号
func (r *TaobaoalitripdomesticrentcarstatusupdateAPIRequest) SetCarNumber(_carNumber string) error {
	r._carNumber = _carNumber
	r.Set("car_number", _carNumber)
	return nil
}

// GetCarNumber CarNumber Getter
func (r TaobaoalitripdomesticrentcarstatusupdateAPIRequest) GetCarNumber() string {
	return r._carNumber
}

// SetExtra is Extra Setter
// JSON扩展值
func (r *TaobaoalitripdomesticrentcarstatusupdateAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r TaobaoalitripdomesticrentcarstatusupdateAPIRequest) GetExtra() string {
	return r._extra
}

// SetStatus is Status Setter
// 121-用车中（用户取车成功） 122-待结算（用户还车成功）
func (r *TaobaoalitripdomesticrentcarstatusupdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoalitripdomesticrentcarstatusupdateAPIRequest) GetStatus() int64 {
	return r._status
}
