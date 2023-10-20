package car

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest 航旅国内租车订单状态更新 API请求
// taobao.alitrip.domestic.rent.car.status.update
//
// 航旅国内租车订单状态更新
type TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest struct {
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

// NewTaobaoAlitripDomesticRentCarStatusUpdateRequest 初始化TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest对象
func NewTaobaoAlitripDomesticRentCarStatusUpdateRequest() *TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest {
	return &TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) Reset() {
	r._thirdOrderId = ""
	r._orderId = ""
	r._providerId = ""
	r._carNumber = ""
	r._extra = ""
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.domestic.rent.car.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetThirdOrderId is ThirdOrderId Setter
// 服务商平台订单号
func (r *TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) SetThirdOrderId(_thirdOrderId string) error {
	r._thirdOrderId = _thirdOrderId
	r.Set("third_order_id", _thirdOrderId)
	return nil
}

// GetThirdOrderId ThirdOrderId Getter
func (r TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) GetThirdOrderId() string {
	return r._thirdOrderId
}

// SetOrderId is OrderId Setter
// 飞猪平台订单号
func (r *TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetProviderId is ProviderId Setter
// 服务商标识，由飞猪提供给到各服务商
func (r *TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) SetProviderId(_providerId string) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) GetProviderId() string {
	return r._providerId
}

// SetCarNumber is CarNumber Setter
// 车牌号
func (r *TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) SetCarNumber(_carNumber string) error {
	r._carNumber = _carNumber
	r.Set("car_number", _carNumber)
	return nil
}

// GetCarNumber CarNumber Getter
func (r TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) GetCarNumber() string {
	return r._carNumber
}

// SetExtra is Extra Setter
// JSON扩展值
func (r *TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) GetExtra() string {
	return r._extra
}

// SetStatus is Status Setter
// 121-用车中（用户取车成功） 122-待结算（用户还车成功）
func (r *TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

var poolTaobaoAlitripDomesticRentCarStatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripDomesticRentCarStatusUpdateRequest()
	},
}

// GetTaobaoAlitripDomesticRentCarStatusUpdateRequest 从 sync.Pool 获取 TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest
func GetTaobaoAlitripDomesticRentCarStatusUpdateAPIRequest() *TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest {
	return poolTaobaoAlitripDomesticRentCarStatusUpdateAPIRequest.Get().(*TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest)
}

// ReleaseTaobaoAlitripDomesticRentCarStatusUpdateAPIRequest 将 TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripDomesticRentCarStatusUpdateAPIRequest(v *TaobaoAlitripDomesticRentCarStatusUpdateAPIRequest) {
	v.Reset()
	poolTaobaoAlitripDomesticRentCarStatusUpdateAPIRequest.Put(v)
}
