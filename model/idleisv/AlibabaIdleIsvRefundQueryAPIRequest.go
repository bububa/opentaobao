package idleisv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvRefundQueryAPIRequest 闲鱼已验货交易订单退款信息查询 API请求
// alibaba.idle.isv.refund.query
//
// 闲鱼服务商交易订单退款信息查询-单个退款查询
type AlibabaIdleIsvRefundQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery
}

// NewAlibabaIdleIsvRefundQueryRequest 初始化AlibabaIdleIsvRefundQueryAPIRequest对象
func NewAlibabaIdleIsvRefundQueryRequest() *AlibabaIdleIsvRefundQueryAPIRequest {
	return &AlibabaIdleIsvRefundQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleIsvRefundQueryAPIRequest) Reset() {
	r._paramAppraiseIsvOrderQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvRefundQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.refund.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvRefundQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvRefundQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAppraiseIsvOrderQuery is ParamAppraiseIsvOrderQuery Setter
// 系统自动生成
func (r *AlibabaIdleIsvRefundQueryAPIRequest) SetParamAppraiseIsvOrderQuery(_paramAppraiseIsvOrderQuery *AppraiseIsvOrderQuery) error {
	r._paramAppraiseIsvOrderQuery = _paramAppraiseIsvOrderQuery
	r.Set("param_appraise_isv_order_query", _paramAppraiseIsvOrderQuery)
	return nil
}

// GetParamAppraiseIsvOrderQuery ParamAppraiseIsvOrderQuery Getter
func (r AlibabaIdleIsvRefundQueryAPIRequest) GetParamAppraiseIsvOrderQuery() *AppraiseIsvOrderQuery {
	return r._paramAppraiseIsvOrderQuery
}

var poolAlibabaIdleIsvRefundQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleIsvRefundQueryRequest()
	},
}

// GetAlibabaIdleIsvRefundQueryRequest 从 sync.Pool 获取 AlibabaIdleIsvRefundQueryAPIRequest
func GetAlibabaIdleIsvRefundQueryAPIRequest() *AlibabaIdleIsvRefundQueryAPIRequest {
	return poolAlibabaIdleIsvRefundQueryAPIRequest.Get().(*AlibabaIdleIsvRefundQueryAPIRequest)
}

// ReleaseAlibabaIdleIsvRefundQueryAPIRequest 将 AlibabaIdleIsvRefundQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleIsvRefundQueryAPIRequest(v *AlibabaIdleIsvRefundQueryAPIRequest) {
	v.Reset()
	poolAlibabaIdleIsvRefundQueryAPIRequest.Put(v)
}
