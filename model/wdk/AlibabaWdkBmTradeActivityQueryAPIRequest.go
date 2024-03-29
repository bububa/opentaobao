package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmTradeActivityQueryAPIRequest 品牌营销的订单活动信息查询 API请求
// alibaba.wdk.bm.trade.activity.query
//
// 品牌营销的订单活动信息查询
type AlibabaWdkBmTradeActivityQueryAPIRequest struct {
	model.Params
	// 入参
	_queryParam *IsvOrderQueryParam
}

// NewAlibabaWdkBmTradeActivityQueryRequest 初始化AlibabaWdkBmTradeActivityQueryAPIRequest对象
func NewAlibabaWdkBmTradeActivityQueryRequest() *AlibabaWdkBmTradeActivityQueryAPIRequest {
	return &AlibabaWdkBmTradeActivityQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkBmTradeActivityQueryAPIRequest) Reset() {
	r._queryParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmTradeActivityQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.trade.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmTradeActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkBmTradeActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryParam is QueryParam Setter
// 入参
func (r *AlibabaWdkBmTradeActivityQueryAPIRequest) SetQueryParam(_queryParam *IsvOrderQueryParam) error {
	r._queryParam = _queryParam
	r.Set("query_param", _queryParam)
	return nil
}

// GetQueryParam QueryParam Getter
func (r AlibabaWdkBmTradeActivityQueryAPIRequest) GetQueryParam() *IsvOrderQueryParam {
	return r._queryParam
}

var poolAlibabaWdkBmTradeActivityQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkBmTradeActivityQueryRequest()
	},
}

// GetAlibabaWdkBmTradeActivityQueryRequest 从 sync.Pool 获取 AlibabaWdkBmTradeActivityQueryAPIRequest
func GetAlibabaWdkBmTradeActivityQueryAPIRequest() *AlibabaWdkBmTradeActivityQueryAPIRequest {
	return poolAlibabaWdkBmTradeActivityQueryAPIRequest.Get().(*AlibabaWdkBmTradeActivityQueryAPIRequest)
}

// ReleaseAlibabaWdkBmTradeActivityQueryAPIRequest 将 AlibabaWdkBmTradeActivityQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkBmTradeActivityQueryAPIRequest(v *AlibabaWdkBmTradeActivityQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkBmTradeActivityQueryAPIRequest.Put(v)
}
