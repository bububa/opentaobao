package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvorderdealrefundAPIRequest 闲鱼无忧购入仓模式服务商退款处理接口 API请求
// alibaba.idle.isv.order.dealrefund
//
// 闲鱼无忧购业务入仓模式下，用户发起退款后，服务商使用此接口处理退款
type AlibabaidleisvorderdealrefundAPIRequest struct {
	model.Params
	// 退款参数
	_paramAppraiseIsvRefundRequest *AppraiseIsvRefundRequest
}

// NewAlibabaidleisvorderdealrefundRequest 初始化AlibabaidleisvorderdealrefundAPIRequest对象
func NewAlibabaidleisvorderdealrefundRequest() *AlibabaidleisvorderdealrefundAPIRequest {
	return &AlibabaidleisvorderdealrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvorderdealrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.order.dealrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvorderdealrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvorderdealrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAppraiseIsvRefundRequest is ParamAppraiseIsvRefundRequest Setter
// 退款参数
func (r *AlibabaidleisvorderdealrefundAPIRequest) SetParamAppraiseIsvRefundRequest(_paramAppraiseIsvRefundRequest *AppraiseIsvRefundRequest) error {
	r._paramAppraiseIsvRefundRequest = _paramAppraiseIsvRefundRequest
	r.Set("param_appraise_isv_refund_request", _paramAppraiseIsvRefundRequest)
	return nil
}

// GetParamAppraiseIsvRefundRequest ParamAppraiseIsvRefundRequest Getter
func (r AlibabaidleisvorderdealrefundAPIRequest) GetParamAppraiseIsvRefundRequest() *AppraiseIsvRefundRequest {
	return r._paramAppraiseIsvRefundRequest
}
