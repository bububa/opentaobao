package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSeriesCreateAPIRequest
系列品变更-新增系列 API请求
alibaba.wdk.series.create

系列品变更-新增系列 */
type AlibabaWdkSeriesCreateAPIRequest struct {
	model.Params
	// 系列品创建系列请求
	_series *SkuSeriesCreateRequest
}

// NewAlibabaWdkSeriesCreateRequest 初始化AlibabaWdkSeriesCreateAPIRequest对象
func NewAlibabaWdkSeriesCreateRequest() *AlibabaWdkSeriesCreateAPIRequest {
	return &AlibabaWdkSeriesCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Series Setter
// 系列品创建系列请求
func (r *AlibabaWdkSeriesCreateAPIRequest) SetSeries(_series *SkuSeriesCreateRequest) error {
	r._series = _series
	r.Set("series", _series)
	return nil
}

// Get Series Getter
func (r AlibabaWdkSeriesCreateAPIRequest) GetSeries() *SkuSeriesCreateRequest {
	return r._series
}
