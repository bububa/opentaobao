package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkserieseditAPIRequest 系列品变更-更新系列 API请求
// alibaba.wdk.series.edit
//
// 系列品变更-更新系列
type AlibabawdkserieseditAPIRequest struct {
	model.Params
	// 商品系列修改请求
	_series *SkuSeriesEditRequest
}

// NewAlibabawdkserieseditRequest 初始化AlibabawdkserieseditAPIRequest对象
func NewAlibabawdkserieseditRequest() *AlibabawdkserieseditAPIRequest {
	return &AlibabawdkserieseditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkserieseditAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkserieseditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkserieseditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeries is Series Setter
// 商品系列修改请求
func (r *AlibabawdkserieseditAPIRequest) SetSeries(_series *SkuSeriesEditRequest) error {
	r._series = _series
	r.Set("series", _series)
	return nil
}

// GetSeries Series Getter
func (r AlibabawdkserieseditAPIRequest) GetSeries() *SkuSeriesEditRequest {
	return r._series
}
