package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkorderlistAPIRequest 五道口订单拉取 API请求
// alibaba.wdk.order.list
//
// 五道口交易订单拉取接口
type AlibabawdkorderlistAPIRequest struct {
	model.Params
	// 查询参数
	_batchQueryRequest *BatchQueryRequest
}

// NewAlibabawdkorderlistRequest 初始化AlibabawdkorderlistAPIRequest对象
func NewAlibabawdkorderlistRequest() *AlibabawdkorderlistAPIRequest {
	return &AlibabawdkorderlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkorderlistAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkorderlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkorderlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchQueryRequest is BatchQueryRequest Setter
// 查询参数
func (r *AlibabawdkorderlistAPIRequest) SetBatchQueryRequest(_batchQueryRequest *BatchQueryRequest) error {
	r._batchQueryRequest = _batchQueryRequest
	r.Set("batch_query_request", _batchQueryRequest)
	return nil
}

// GetBatchQueryRequest BatchQueryRequest Getter
func (r AlibabawdkorderlistAPIRequest) GetBatchQueryRequest() *BatchQueryRequest {
	return r._batchQueryRequest
}
