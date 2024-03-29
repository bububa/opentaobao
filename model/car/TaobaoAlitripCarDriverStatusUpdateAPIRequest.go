package car

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarDriverStatusUpdateAPIRequest 司机服务状态更新接口 API请求
// taobao.alitrip.car.driver.status.update
//
// 飞猪用车业务回调接口，用于服务商实时回传更新司机当前服务状态
type TaobaoAlitripCarDriverStatusUpdateAPIRequest struct {
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

// NewTaobaoAlitripCarDriverStatusUpdateRequest 初始化TaobaoAlitripCarDriverStatusUpdateAPIRequest对象
func NewTaobaoAlitripCarDriverStatusUpdateRequest() *TaobaoAlitripCarDriverStatusUpdateAPIRequest {
	return &TaobaoAlitripCarDriverStatusUpdateAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripCarDriverStatusUpdateAPIRequest) Reset() {
	r._orderId = ""
	r._thirdOrderId = ""
	r._providerId = ""
	r._time = ""
	r._sellerId = ""
	r._status = 0
	r._useType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarDriverStatusUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.driver.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarDriverStatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripCarDriverStatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 飞猪订单id
func (r *TaobaoAlitripCarDriverStatusUpdateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoAlitripCarDriverStatusUpdateAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetThirdOrderId is ThirdOrderId Setter
// 服务商订单id
func (r *TaobaoAlitripCarDriverStatusUpdateAPIRequest) SetThirdOrderId(_thirdOrderId string) error {
	r._thirdOrderId = _thirdOrderId
	r.Set("third_order_id", _thirdOrderId)
	return nil
}

// GetThirdOrderId ThirdOrderId Getter
func (r TaobaoAlitripCarDriverStatusUpdateAPIRequest) GetThirdOrderId() string {
	return r._thirdOrderId
}

// SetProviderId is ProviderId Setter
// 服务商标识，由飞猪提供给到各服务商
func (r *TaobaoAlitripCarDriverStatusUpdateAPIRequest) SetProviderId(_providerId string) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r TaobaoAlitripCarDriverStatusUpdateAPIRequest) GetProviderId() string {
	return r._providerId
}

// SetTime is Time Setter
// 状态变更相应时间（如司机出发时间、司机到达时间、服务开始时间），格式：yyyy-mm-dd hh:mm:ss
func (r *TaobaoAlitripCarDriverStatusUpdateAPIRequest) SetTime(_time string) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// GetTime Time Getter
func (r TaobaoAlitripCarDriverStatusUpdateAPIRequest) GetTime() string {
	return r._time
}

// SetSellerId is SellerId Setter
// 可选，卖家id
func (r *TaobaoAlitripCarDriverStatusUpdateAPIRequest) SetSellerId(_sellerId string) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TaobaoAlitripCarDriverStatusUpdateAPIRequest) GetSellerId() string {
	return r._sellerId
}

// SetStatus is Status Setter
// 司机服务状态。1-司机已出发，2-司机已到达，3-司机已开始服务
func (r *TaobaoAlitripCarDriverStatusUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoAlitripCarDriverStatusUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetUseType is UseType Setter
// 0:接送机 1：实时打车 2：租车(不传值默认为0)
func (r *TaobaoAlitripCarDriverStatusUpdateAPIRequest) SetUseType(_useType int64) error {
	r._useType = _useType
	r.Set("use_type", _useType)
	return nil
}

// GetUseType UseType Getter
func (r TaobaoAlitripCarDriverStatusUpdateAPIRequest) GetUseType() int64 {
	return r._useType
}

var poolTaobaoAlitripCarDriverStatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripCarDriverStatusUpdateRequest()
	},
}

// GetTaobaoAlitripCarDriverStatusUpdateRequest 从 sync.Pool 获取 TaobaoAlitripCarDriverStatusUpdateAPIRequest
func GetTaobaoAlitripCarDriverStatusUpdateAPIRequest() *TaobaoAlitripCarDriverStatusUpdateAPIRequest {
	return poolTaobaoAlitripCarDriverStatusUpdateAPIRequest.Get().(*TaobaoAlitripCarDriverStatusUpdateAPIRequest)
}

// ReleaseTaobaoAlitripCarDriverStatusUpdateAPIRequest 将 TaobaoAlitripCarDriverStatusUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripCarDriverStatusUpdateAPIRequest(v *TaobaoAlitripCarDriverStatusUpdateAPIRequest) {
	v.Reset()
	poolTaobaoAlitripCarDriverStatusUpdateAPIRequest.Put(v)
}
