package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvOrderDealrefundAPIRequest 闲鱼无忧购入仓模式服务商退款处理接口 API请求
// alibaba.idle.isv.order.dealrefund
//
// 闲鱼无忧购业务入仓模式下，用户发起退款后，服务商使用此接口处理退款
type AlibabaIdleIsvOrderDealrefundAPIRequest struct {
	model.Params
	// 退款参数
	_paramAppraiseIsvRefundRequest *AppraiseIsvRefundRequest
}

// NewAlibabaIdleIsvOrderDealrefundRequest 初始化AlibabaIdleIsvOrderDealrefundAPIRequest对象
func NewAlibabaIdleIsvOrderDealrefundRequest() *AlibabaIdleIsvOrderDealrefundAPIRequest {
	return &AlibabaIdleIsvOrderDealrefundAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleIsvOrderDealrefundAPIRequest) Reset() {
	r._paramAppraiseIsvRefundRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderDealrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.order.dealrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderDealrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvOrderDealrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAppraiseIsvRefundRequest is ParamAppraiseIsvRefundRequest Setter
// 退款参数
func (r *AlibabaIdleIsvOrderDealrefundAPIRequest) SetParamAppraiseIsvRefundRequest(_paramAppraiseIsvRefundRequest *AppraiseIsvRefundRequest) error {
	r._paramAppraiseIsvRefundRequest = _paramAppraiseIsvRefundRequest
	r.Set("param_appraise_isv_refund_request", _paramAppraiseIsvRefundRequest)
	return nil
}

// GetParamAppraiseIsvRefundRequest ParamAppraiseIsvRefundRequest Getter
func (r AlibabaIdleIsvOrderDealrefundAPIRequest) GetParamAppraiseIsvRefundRequest() *AppraiseIsvRefundRequest {
	return r._paramAppraiseIsvRefundRequest
}

var poolAlibabaIdleIsvOrderDealrefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleIsvOrderDealrefundRequest()
	},
}

// GetAlibabaIdleIsvOrderDealrefundRequest 从 sync.Pool 获取 AlibabaIdleIsvOrderDealrefundAPIRequest
func GetAlibabaIdleIsvOrderDealrefundAPIRequest() *AlibabaIdleIsvOrderDealrefundAPIRequest {
	return poolAlibabaIdleIsvOrderDealrefundAPIRequest.Get().(*AlibabaIdleIsvOrderDealrefundAPIRequest)
}

// ReleaseAlibabaIdleIsvOrderDealrefundAPIRequest 将 AlibabaIdleIsvOrderDealrefundAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleIsvOrderDealrefundAPIRequest(v *AlibabaIdleIsvOrderDealrefundAPIRequest) {
	v.Reset()
	poolAlibabaIdleIsvOrderDealrefundAPIRequest.Put(v)
}
