package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkseriesskuaddAPIRequest 系列品商品变更-添加商品 API请求
// alibaba.wdk.series.sku.add
//
// 系列品商品变更-添加商品
type AlibabawdkseriesskuaddAPIRequest struct {
	model.Params
	// 系列品添加商品请求
	_seriesSkus *SeriesSkuRequest
}

// NewAlibabawdkseriesskuaddRequest 初始化AlibabawdkseriesskuaddAPIRequest对象
func NewAlibabawdkseriesskuaddRequest() *AlibabawdkseriesskuaddAPIRequest {
	return &AlibabawdkseriesskuaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkseriesskuaddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkseriesskuaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkseriesskuaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeriesSkus is SeriesSkus Setter
// 系列品添加商品请求
func (r *AlibabawdkseriesskuaddAPIRequest) SetSeriesSkus(_seriesSkus *SeriesSkuRequest) error {
	r._seriesSkus = _seriesSkus
	r.Set("series_skus", _seriesSkus)
	return nil
}

// GetSeriesSkus SeriesSkus Getter
func (r AlibabawdkseriesskuaddAPIRequest) GetSeriesSkus() *SeriesSkuRequest {
	return r._seriesSkus
}
