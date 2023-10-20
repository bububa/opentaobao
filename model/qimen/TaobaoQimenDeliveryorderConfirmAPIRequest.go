package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderConfirmAPIRequest 发货单确认接口 API请求
// taobao.qimen.deliveryorder.confirm
//
// taobao.qimen.deliveryorder.confirm
type TaobaoQimenDeliveryorderConfirmAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderConfirmRequest
}

// NewTaobaoQimenDeliveryorderConfirmRequest 初始化TaobaoQimenDeliveryorderConfirmAPIRequest对象
func NewTaobaoQimenDeliveryorderConfirmRequest() *TaobaoQimenDeliveryorderConfirmAPIRequest {
	return &TaobaoQimenDeliveryorderConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenDeliveryorderConfirmAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenDeliveryorderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenDeliveryorderConfirmAPIRequest) SetRequest(_request *DeliveryOrderConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenDeliveryorderConfirmAPIRequest) GetRequest() *DeliveryOrderConfirmRequest {
	return r._request
}

var poolTaobaoQimenDeliveryorderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenDeliveryorderConfirmRequest()
	},
}

// GetTaobaoQimenDeliveryorderConfirmRequest 从 sync.Pool 获取 TaobaoQimenDeliveryorderConfirmAPIRequest
func GetTaobaoQimenDeliveryorderConfirmAPIRequest() *TaobaoQimenDeliveryorderConfirmAPIRequest {
	return poolTaobaoQimenDeliveryorderConfirmAPIRequest.Get().(*TaobaoQimenDeliveryorderConfirmAPIRequest)
}

// ReleaseTaobaoQimenDeliveryorderConfirmAPIRequest 将 TaobaoQimenDeliveryorderConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenDeliveryorderConfirmAPIRequest(v *TaobaoQimenDeliveryorderConfirmAPIRequest) {
	v.Reset()
	poolTaobaoQimenDeliveryorderConfirmAPIRequest.Put(v)
}
