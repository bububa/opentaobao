package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunaliospaytradequeryAPIRequest 查询支付结果接口 API请求
// aliyun.alios.pay.trade.query
//
// 商户用来查询支付结果接口
type AliyunaliospaytradequeryAPIRequest struct {
	model.Params
	// 请求参数
	_queryTradeRequest *QueryTradeRequest
}

// NewAliyunaliospaytradequeryRequest 初始化AliyunaliospaytradequeryAPIRequest对象
func NewAliyunaliospaytradequeryRequest() *AliyunaliospaytradequeryAPIRequest {
	return &AliyunaliospaytradequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunaliospaytradequeryAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.trade.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunaliospaytradequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunaliospaytradequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryTradeRequest is QueryTradeRequest Setter
// 请求参数
func (r *AliyunaliospaytradequeryAPIRequest) SetQueryTradeRequest(_queryTradeRequest *QueryTradeRequest) error {
	r._queryTradeRequest = _queryTradeRequest
	r.Set("query_trade_request", _queryTradeRequest)
	return nil
}

// GetQueryTradeRequest QueryTradeRequest Getter
func (r AliyunaliospaytradequeryAPIRequest) GetQueryTradeRequest() *QueryTradeRequest {
	return r._queryTradeRequest
}
