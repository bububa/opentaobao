package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesEditAPIRequest 系列品变更-更新系列 API请求
// alibaba.wdk.series.edit
//
// 系列品变更-更新系列
type AlibabaWdkSeriesEditAPIRequest struct {
	model.Params
	// 商品系列修改请求
	_series *SkuSeriesEditRequest
}

// NewAlibabaWdkSeriesEditRequest 初始化AlibabaWdkSeriesEditAPIRequest对象
func NewAlibabaWdkSeriesEditRequest() *AlibabaWdkSeriesEditAPIRequest {
	return &AlibabaWdkSeriesEditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesEditAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesEditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Series Setter
// 商品系列修改请求
func (r *AlibabaWdkSeriesEditAPIRequest) SetSeries(_series *SkuSeriesEditRequest) error {
	r._series = _series
	r.Set("series", _series)
	return nil
}

// Get Series Getter
func (r AlibabaWdkSeriesEditAPIRequest) GetSeries() *SkuSeriesEditRequest {
	return r._series
}
