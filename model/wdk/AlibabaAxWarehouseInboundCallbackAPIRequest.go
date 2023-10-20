package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAxWarehouseInboundCallbackAPIRequest 翱象入库回传 API请求
// alibaba.ax.warehouse.inbound.callback
//
// 翱象入库回传
type AlibabaAxWarehouseInboundCallbackAPIRequest struct {
	model.Params
	// 消息体
	_reverseInBoundCallBackRequest *TopReverseInBoundCallBackRequest
}

// NewAlibabaAxWarehouseInboundCallbackRequest 初始化AlibabaAxWarehouseInboundCallbackAPIRequest对象
func NewAlibabaAxWarehouseInboundCallbackRequest() *AlibabaAxWarehouseInboundCallbackAPIRequest {
	return &AlibabaAxWarehouseInboundCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAxWarehouseInboundCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ax.warehouse.inbound.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAxWarehouseInboundCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAxWarehouseInboundCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReverseInBoundCallBackRequest is ReverseInBoundCallBackRequest Setter
// 消息体
func (r *AlibabaAxWarehouseInboundCallbackAPIRequest) SetReverseInBoundCallBackRequest(_reverseInBoundCallBackRequest *TopReverseInBoundCallBackRequest) error {
	r._reverseInBoundCallBackRequest = _reverseInBoundCallBackRequest
	r.Set("reverse_in_bound_call_back_request", _reverseInBoundCallBackRequest)
	return nil
}

// GetReverseInBoundCallBackRequest ReverseInBoundCallBackRequest Getter
func (r AlibabaAxWarehouseInboundCallbackAPIRequest) GetReverseInBoundCallBackRequest() *TopReverseInBoundCallBackRequest {
	return r._reverseInBoundCallBackRequest
}
