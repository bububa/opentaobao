package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesSortAPIRequest 系列品-商品排序 API请求
// alibaba.wdk.series.sort
//
// 系列品商品变更-商品排序
type AlibabaWdkSeriesSortAPIRequest struct {
	model.Params
	// 自定义排序请求
	_sort *SeriesSortRequest
}

// NewAlibabaWdkSeriesSortRequest 初始化AlibabaWdkSeriesSortAPIRequest对象
func NewAlibabaWdkSeriesSortRequest() *AlibabaWdkSeriesSortAPIRequest {
	return &AlibabaWdkSeriesSortAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesSortAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.sort"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesSortAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSeriesSortAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSort is Sort Setter
// 自定义排序请求
func (r *AlibabaWdkSeriesSortAPIRequest) SetSort(_sort *SeriesSortRequest) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r AlibabaWdkSeriesSortAPIRequest) GetSort() *SeriesSortRequest {
	return r._sort
}
