package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderDelivererChangeAPIRequest 配送员信息变更接口 API请求
// alibaba.aelophy.order.deliverer.change
//
// 配送员信息变更接口
type AlibabaAelophyOrderDelivererChangeAPIRequest struct {
	model.Params
	// 配送员信息变更请求
	_delivererChangeRequest *DelivererChangeRequest
}

// NewAlibabaAelophyOrderDelivererChangeRequest 初始化AlibabaAelophyOrderDelivererChangeAPIRequest对象
func NewAlibabaAelophyOrderDelivererChangeRequest() *AlibabaAelophyOrderDelivererChangeAPIRequest {
	return &AlibabaAelophyOrderDelivererChangeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAelophyOrderDelivererChangeAPIRequest) Reset() {
	r._delivererChangeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAelophyOrderDelivererChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.order.deliverer.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAelophyOrderDelivererChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAelophyOrderDelivererChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDelivererChangeRequest is DelivererChangeRequest Setter
// 配送员信息变更请求
func (r *AlibabaAelophyOrderDelivererChangeAPIRequest) SetDelivererChangeRequest(_delivererChangeRequest *DelivererChangeRequest) error {
	r._delivererChangeRequest = _delivererChangeRequest
	r.Set("deliverer_change_request", _delivererChangeRequest)
	return nil
}

// GetDelivererChangeRequest DelivererChangeRequest Getter
func (r AlibabaAelophyOrderDelivererChangeAPIRequest) GetDelivererChangeRequest() *DelivererChangeRequest {
	return r._delivererChangeRequest
}

var poolAlibabaAelophyOrderDelivererChangeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAelophyOrderDelivererChangeRequest()
	},
}

// GetAlibabaAelophyOrderDelivererChangeRequest 从 sync.Pool 获取 AlibabaAelophyOrderDelivererChangeAPIRequest
func GetAlibabaAelophyOrderDelivererChangeAPIRequest() *AlibabaAelophyOrderDelivererChangeAPIRequest {
	return poolAlibabaAelophyOrderDelivererChangeAPIRequest.Get().(*AlibabaAelophyOrderDelivererChangeAPIRequest)
}

// ReleaseAlibabaAelophyOrderDelivererChangeAPIRequest 将 AlibabaAelophyOrderDelivererChangeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAelophyOrderDelivererChangeAPIRequest(v *AlibabaAelophyOrderDelivererChangeAPIRequest) {
	v.Reset()
	poolAlibabaAelophyOrderDelivererChangeAPIRequest.Put(v)
}
