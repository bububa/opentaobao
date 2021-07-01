package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvOrderDealrefundAPIRequest
闲鱼无忧购入仓模式服务商退款处理接口 API请求
alibaba.idle.isv.order.dealrefund

闲鱼无忧购业务入仓模式下，用户发起退款后，服务商使用此接口处理退款 */
type AlibabaIdleIsvOrderDealrefundAPIRequest struct {
	model.Params
	// 退款参数
	_paramAppraiseIsvRefundRequest *AppraiseIsvRefundRequest
}

// NewAlibabaIdleIsvOrderDealrefundRequest 初始化AlibabaIdleIsvOrderDealrefundAPIRequest对象
func NewAlibabaIdleIsvOrderDealrefundRequest() *AlibabaIdleIsvOrderDealrefundAPIRequest {
	return &AlibabaIdleIsvOrderDealrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderDealrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.order.dealrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderDealrefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamAppraiseIsvRefundRequest Setter
// 退款参数
func (r *AlibabaIdleIsvOrderDealrefundAPIRequest) SetParamAppraiseIsvRefundRequest(_paramAppraiseIsvRefundRequest *AppraiseIsvRefundRequest) error {
	r._paramAppraiseIsvRefundRequest = _paramAppraiseIsvRefundRequest
	r.Set("param_appraise_isv_refund_request", _paramAppraiseIsvRefundRequest)
	return nil
}

// Get ParamAppraiseIsvRefundRequest Getter
func (r AlibabaIdleIsvOrderDealrefundAPIRequest) GetParamAppraiseIsvRefundRequest() *AppraiseIsvRefundRequest {
	return r._paramAppraiseIsvRefundRequest
}
