package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkseriessortAPIRequest 系列品-商品排序 API请求
// alibaba.wdk.series.sort
//
// 系列品商品变更-商品排序
type AlibabawdkseriessortAPIRequest struct {
	model.Params
	// 自定义排序请求
	_sort *SeriesSortRequest
}

// NewAlibabawdkseriessortRequest 初始化AlibabawdkseriessortAPIRequest对象
func NewAlibabawdkseriessortRequest() *AlibabawdkseriessortAPIRequest {
	return &AlibabawdkseriessortAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkseriessortAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.sort"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkseriessortAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkseriessortAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSort is Sort Setter
// 自定义排序请求
func (r *AlibabawdkseriessortAPIRequest) SetSort(_sort *SeriesSortRequest) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r AlibabawdkseriessortAPIRequest) GetSort() *SeriesSortRequest {
	return r._sort
}
