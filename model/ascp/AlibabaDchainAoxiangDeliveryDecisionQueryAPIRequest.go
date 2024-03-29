package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest 查询黑白名单快递 API请求
// alibaba.dchain.aoxiang.delivery.decision.query
//
// 查询黑白名单快递
type AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest struct {
	model.Params
	// 黑白名单快递查询入参
	_deliveryDecisionQueryRequest *DeliveryDecisionQueryRequest
}

// NewAlibabaDchainAoxiangDeliveryDecisionQueryRequest 初始化AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest对象
func NewAlibabaDchainAoxiangDeliveryDecisionQueryRequest() *AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest {
	return &AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest) Reset() {
	r._deliveryDecisionQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.delivery.decision.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryDecisionQueryRequest is DeliveryDecisionQueryRequest Setter
// 黑白名单快递查询入参
func (r *AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest) SetDeliveryDecisionQueryRequest(_deliveryDecisionQueryRequest *DeliveryDecisionQueryRequest) error {
	r._deliveryDecisionQueryRequest = _deliveryDecisionQueryRequest
	r.Set("delivery_decision_query_request", _deliveryDecisionQueryRequest)
	return nil
}

// GetDeliveryDecisionQueryRequest DeliveryDecisionQueryRequest Getter
func (r AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest) GetDeliveryDecisionQueryRequest() *DeliveryDecisionQueryRequest {
	return r._deliveryDecisionQueryRequest
}

var poolAlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangDeliveryDecisionQueryRequest()
	},
}

// GetAlibabaDchainAoxiangDeliveryDecisionQueryRequest 从 sync.Pool 获取 AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest
func GetAlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest() *AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest {
	return poolAlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest.Get().(*AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest)
}

// ReleaseAlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest 将 AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest(v *AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest.Put(v)
}
