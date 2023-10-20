package ascpqcc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpqccsamplecancelitemrelationAPIRequest 魅力惠样品解除父子商品关系 API请求
// alibaba.ascp.qcc.sample.cancel.item.relation
//
// 品控中心魅力惠样品解除父子商品关系
type AlibabaascpqccsamplecancelitemrelationAPIRequest struct {
	model.Params
	// 请求参数对象
	_cancelRequest *CancelSampleRelationRequest
}

// NewAlibabaascpqccsamplecancelitemrelationRequest 初始化AlibabaascpqccsamplecancelitemrelationAPIRequest对象
func NewAlibabaascpqccsamplecancelitemrelationRequest() *AlibabaascpqccsamplecancelitemrelationAPIRequest {
	return &AlibabaascpqccsamplecancelitemrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpqccsamplecancelitemrelationAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.qcc.sample.cancel.item.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpqccsamplecancelitemrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpqccsamplecancelitemrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelRequest is CancelRequest Setter
// 请求参数对象
func (r *AlibabaascpqccsamplecancelitemrelationAPIRequest) SetCancelRequest(_cancelRequest *CancelSampleRelationRequest) error {
	r._cancelRequest = _cancelRequest
	r.Set("cancel_request", _cancelRequest)
	return nil
}

// GetCancelRequest CancelRequest Getter
func (r AlibabaascpqccsamplecancelitemrelationAPIRequest) GetCancelRequest() *CancelSampleRelationRequest {
	return r._cancelRequest
}
