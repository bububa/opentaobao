package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripcardriverstatusupdateAPIRequest 司机服务状态更新接口 API请求
// taobao.alitrip.car.driver.status.update
//
// 飞猪用车业务回调接口，用于服务商实时回传更新司机当前服务状态
type TaobaoalitripcardriverstatusupdateAPIRequest struct {
	model.Params
	// 飞猪订单id
	_orderId string
	// 服务商订单id
	_thirdOrderId string
	// 服务商标识，由飞猪提供给到各服务商
	_providerId string
	// 状态变更相应时间（如司机出发时间、司机到达时间、服务开始时间），格式：yyyy-mm-dd hh:mm:ss
	_time string
	// 可选，卖家id
	_sellerId string
	// 司机服务状态。1-司机已出发，2-司机已到达，3-司机已开始服务
	_status int64
	// 0:接送机 1：实时打车 2：租车(不传值默认为0)
	_useType int64
}

// NewTaobaoalitripcardriverstatusupdateRequest 初始化TaobaoalitripcardriverstatusupdateAPIRequest对象
func NewTaobaoalitripcardriverstatusupdateRequest() *TaobaoalitripcardriverstatusupdateAPIRequest {
	return &TaobaoalitripcardriverstatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripcardriverstatusupdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.driver.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripcardriverstatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripcardriverstatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 飞猪订单id
func (r *TaobaoalitripcardriverstatusupdateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoalitripcardriverstatusupdateAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetThirdOrderId is ThirdOrderId Setter
// 服务商订单id
func (r *TaobaoalitripcardriverstatusupdateAPIRequest) SetThirdOrderId(_thirdOrderId string) error {
	r._thirdOrderId = _thirdOrderId
	r.Set("third_order_id", _thirdOrderId)
	return nil
}

// GetThirdOrderId ThirdOrderId Getter
func (r TaobaoalitripcardriverstatusupdateAPIRequest) GetThirdOrderId() string {
	return r._thirdOrderId
}

// SetProviderId is ProviderId Setter
// 服务商标识，由飞猪提供给到各服务商
func (r *TaobaoalitripcardriverstatusupdateAPIRequest) SetProviderId(_providerId string) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r TaobaoalitripcardriverstatusupdateAPIRequest) GetProviderId() string {
	return r._providerId
}

// SetTime is Time Setter
// 状态变更相应时间（如司机出发时间、司机到达时间、服务开始时间），格式：yyyy-mm-dd hh:mm:ss
func (r *TaobaoalitripcardriverstatusupdateAPIRequest) SetTime(_time string) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// GetTime Time Getter
func (r TaobaoalitripcardriverstatusupdateAPIRequest) GetTime() string {
	return r._time
}

// SetSellerId is SellerId Setter
// 可选，卖家id
func (r *TaobaoalitripcardriverstatusupdateAPIRequest) SetSellerId(_sellerId string) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TaobaoalitripcardriverstatusupdateAPIRequest) GetSellerId() string {
	return r._sellerId
}

// SetStatus is Status Setter
// 司机服务状态。1-司机已出发，2-司机已到达，3-司机已开始服务
func (r *TaobaoalitripcardriverstatusupdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoalitripcardriverstatusupdateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetUseType is UseType Setter
// 0:接送机 1：实时打车 2：租车(不传值默认为0)
func (r *TaobaoalitripcardriverstatusupdateAPIRequest) SetUseType(_useType int64) error {
	r._useType = _useType
	r.Set("use_type", _useType)
	return nil
}

// GetUseType UseType Getter
func (r TaobaoalitripcardriverstatusupdateAPIRequest) GetUseType() int64 {
	return r._useType
}
