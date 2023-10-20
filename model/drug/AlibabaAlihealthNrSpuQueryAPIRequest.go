package drug

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrSpuQueryAPIRequest 获取标品库标品信息 API请求
// alibaba.alihealth.nr.spu.query
//
// 提供给ERP使用的，获取健康标品库信息
type AlibabaAlihealthNrSpuQueryAPIRequest struct {
	model.Params
	// 标品查询条件
	_query *TopAlihealthSpuQuery
	// 查询选择器
	_options *TopAlihealthSpuQueryOptions
}

// NewAlibabaAlihealthNrSpuQueryRequest 初始化AlibabaAlihealthNrSpuQueryAPIRequest对象
func NewAlibabaAlihealthNrSpuQueryRequest() *AlibabaAlihealthNrSpuQueryAPIRequest {
	return &AlibabaAlihealthNrSpuQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthNrSpuQueryAPIRequest) Reset() {
	r._query = nil
	r._options = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrSpuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.spu.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrSpuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthNrSpuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 标品查询条件
func (r *AlibabaAlihealthNrSpuQueryAPIRequest) SetQuery(_query *TopAlihealthSpuQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaAlihealthNrSpuQueryAPIRequest) GetQuery() *TopAlihealthSpuQuery {
	return r._query
}

// SetOptions is Options Setter
// 查询选择器
func (r *AlibabaAlihealthNrSpuQueryAPIRequest) SetOptions(_options *TopAlihealthSpuQueryOptions) error {
	r._options = _options
	r.Set("options", _options)
	return nil
}

// GetOptions Options Getter
func (r AlibabaAlihealthNrSpuQueryAPIRequest) GetOptions() *TopAlihealthSpuQueryOptions {
	return r._options
}

var poolAlibabaAlihealthNrSpuQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthNrSpuQueryRequest()
	},
}

// GetAlibabaAlihealthNrSpuQueryRequest 从 sync.Pool 获取 AlibabaAlihealthNrSpuQueryAPIRequest
func GetAlibabaAlihealthNrSpuQueryAPIRequest() *AlibabaAlihealthNrSpuQueryAPIRequest {
	return poolAlibabaAlihealthNrSpuQueryAPIRequest.Get().(*AlibabaAlihealthNrSpuQueryAPIRequest)
}

// ReleaseAlibabaAlihealthNrSpuQueryAPIRequest 将 AlibabaAlihealthNrSpuQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthNrSpuQueryAPIRequest(v *AlibabaAlihealthNrSpuQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthNrSpuQueryAPIRequest.Put(v)
}
