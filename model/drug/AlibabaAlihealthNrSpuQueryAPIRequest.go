package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthnrspuqueryAPIRequest 获取标品库标品信息 API请求
// alibaba.alihealth.nr.spu.query
//
// 提供给ERP使用的，获取健康标品库信息
type AlibabaalihealthnrspuqueryAPIRequest struct {
	model.Params
	// 标品查询条件
	_query *TopAlihealthSpuQuery
	// 查询选择器
	_options *TopAlihealthSpuQueryOptions
}

// NewAlibabaalihealthnrspuqueryRequest 初始化AlibabaalihealthnrspuqueryAPIRequest对象
func NewAlibabaalihealthnrspuqueryRequest() *AlibabaalihealthnrspuqueryAPIRequest {
	return &AlibabaalihealthnrspuqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthnrspuqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.spu.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthnrspuqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthnrspuqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 标品查询条件
func (r *AlibabaalihealthnrspuqueryAPIRequest) SetQuery(_query *TopAlihealthSpuQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaalihealthnrspuqueryAPIRequest) GetQuery() *TopAlihealthSpuQuery {
	return r._query
}

// SetOptions is Options Setter
// 查询选择器
func (r *AlibabaalihealthnrspuqueryAPIRequest) SetOptions(_options *TopAlihealthSpuQueryOptions) error {
	r._options = _options
	r.Set("options", _options)
	return nil
}

// GetOptions Options Getter
func (r AlibabaalihealthnrspuqueryAPIRequest) GetOptions() *TopAlihealthSpuQueryOptions {
	return r._options
}
