package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofuwupurchaseorderpayAPIRequest 内购服务订单付款页获取接口 API请求
// taobao.fuwu.purchase.order.pay
//
// 通过接口获取某一订单的付款页面链接
type TaobaofuwupurchaseorderpayAPIRequest struct {
	model.Params
	// APPKEY，必填
	_appkey string
	// 设备类型，目前只支持PC，可选
	_deviceType string
	// 外部订单号，使用该参数完成查询订单等操作，与外部订单号二选一
	_outOrderId string
	// 订单号，与外部订单号二选一
	_orderId int64
}

// NewTaobaofuwupurchaseorderpayRequest 初始化TaobaofuwupurchaseorderpayAPIRequest对象
func NewTaobaofuwupurchaseorderpayRequest() *TaobaofuwupurchaseorderpayAPIRequest {
	return &TaobaofuwupurchaseorderpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofuwupurchaseorderpayAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.purchase.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofuwupurchaseorderpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofuwupurchaseorderpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppkey is Appkey Setter
// APPKEY，必填
func (r *TaobaofuwupurchaseorderpayAPIRequest) SetAppkey(_appkey string) error {
	r._appkey = _appkey
	r.Set("appkey", _appkey)
	return nil
}

// GetAppkey Appkey Getter
func (r TaobaofuwupurchaseorderpayAPIRequest) GetAppkey() string {
	return r._appkey
}

// SetDeviceType is DeviceType Setter
// 设备类型，目前只支持PC，可选
func (r *TaobaofuwupurchaseorderpayAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaofuwupurchaseorderpayAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetOutOrderId is OutOrderId Setter
// 外部订单号，使用该参数完成查询订单等操作，与外部订单号二选一
func (r *TaobaofuwupurchaseorderpayAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r TaobaofuwupurchaseorderpayAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetOrderId is OrderId Setter
// 订单号，与外部订单号二选一
func (r *TaobaofuwupurchaseorderpayAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaofuwupurchaseorderpayAPIRequest) GetOrderId() int64 {
	return r._orderId
}
