package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesSkuRemoveAPIRequest 系列品商品变更-移除商品 API请求
// alibaba.wdk.series.sku.remove
//
// 系列品商品变更-移除商品
type AlibabaWdkSeriesSkuRemoveAPIRequest struct {
	model.Params
	// 系列品移除商品请求
	_seriesSkus *SeriesSkuRequest
}

// NewAlibabaWdkSeriesSkuRemoveRequest 初始化AlibabaWdkSeriesSkuRemoveAPIRequest对象
func NewAlibabaWdkSeriesSkuRemoveRequest() *AlibabaWdkSeriesSkuRemoveAPIRequest {
	return &AlibabaWdkSeriesSkuRemoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesSkuRemoveAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.sku.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesSkuRemoveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SeriesSkus Setter
// 系列品移除商品请求
func (r *AlibabaWdkSeriesSkuRemoveAPIRequest) SetSeriesSkus(_seriesSkus *SeriesSkuRequest) error {
	r._seriesSkus = _seriesSkus
	r.Set("series_skus", _seriesSkus)
	return nil
}

// Get SeriesSkus Getter
func (r AlibabaWdkSeriesSkuRemoveAPIRequest) GetSeriesSkus() *SeriesSkuRequest {
	return r._seriesSkus
}
