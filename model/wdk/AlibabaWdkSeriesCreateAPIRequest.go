package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkseriescreateAPIRequest 系列品变更-新增系列 API请求
// alibaba.wdk.series.create
//
// 系列品变更-新增系列
type AlibabawdkseriescreateAPIRequest struct {
	model.Params
	// 系列品创建系列请求
	_series *SkuSeriesCreateRequest
}

// NewAlibabawdkseriescreateRequest 初始化AlibabawdkseriescreateAPIRequest对象
func NewAlibabawdkseriescreateRequest() *AlibabawdkseriescreateAPIRequest {
	return &AlibabawdkseriescreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkseriescreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkseriescreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkseriescreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeries is Series Setter
// 系列品创建系列请求
func (r *AlibabawdkseriescreateAPIRequest) SetSeries(_series *SkuSeriesCreateRequest) error {
	r._series = _series
	r.Set("series", _series)
	return nil
}

// GetSeries Series Getter
func (r AlibabawdkseriescreateAPIRequest) GetSeries() *SkuSeriesCreateRequest {
	return r._series
}
