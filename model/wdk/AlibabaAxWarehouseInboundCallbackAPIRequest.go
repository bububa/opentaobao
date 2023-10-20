package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaxwarehouseinboundcallbackAPIRequest 翱象入库回传 API请求
// alibaba.ax.warehouse.inbound.callback
//
// 翱象入库回传
type AlibabaaxwarehouseinboundcallbackAPIRequest struct {
	model.Params
	// 消息体
	_reverseInBoundCallBackRequest *TopReverseInBoundCallBackRequest
}

// NewAlibabaaxwarehouseinboundcallbackRequest 初始化AlibabaaxwarehouseinboundcallbackAPIRequest对象
func NewAlibabaaxwarehouseinboundcallbackRequest() *AlibabaaxwarehouseinboundcallbackAPIRequest {
	return &AlibabaaxwarehouseinboundcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaxwarehouseinboundcallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.ax.warehouse.inbound.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaxwarehouseinboundcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaxwarehouseinboundcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReverseInBoundCallBackRequest is ReverseInBoundCallBackRequest Setter
// 消息体
func (r *AlibabaaxwarehouseinboundcallbackAPIRequest) SetReverseInBoundCallBackRequest(_reverseInBoundCallBackRequest *TopReverseInBoundCallBackRequest) error {
	r._reverseInBoundCallBackRequest = _reverseInBoundCallBackRequest
	r.Set("reverse_in_bound_call_back_request", _reverseInBoundCallBackRequest)
	return nil
}

// GetReverseInBoundCallBackRequest ReverseInBoundCallBackRequest Getter
func (r AlibabaaxwarehouseinboundcallbackAPIRequest) GetReverseInBoundCallBackRequest() *TopReverseInBoundCallBackRequest {
	return r._reverseInBoundCallBackRequest
}
