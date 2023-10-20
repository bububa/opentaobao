package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkseriesskuremoveAPIRequest 系列品商品变更-移除商品 API请求
// alibaba.wdk.series.sku.remove
//
// 系列品商品变更-移除商品
type AlibabawdkseriesskuremoveAPIRequest struct {
	model.Params
	// 系列品移除商品请求
	_seriesSkus *SeriesSkuRequest
}

// NewAlibabawdkseriesskuremoveRequest 初始化AlibabawdkseriesskuremoveAPIRequest对象
func NewAlibabawdkseriesskuremoveRequest() *AlibabawdkseriesskuremoveAPIRequest {
	return &AlibabawdkseriesskuremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkseriesskuremoveAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.sku.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkseriesskuremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkseriesskuremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeriesSkus is SeriesSkus Setter
// 系列品移除商品请求
func (r *AlibabawdkseriesskuremoveAPIRequest) SetSeriesSkus(_seriesSkus *SeriesSkuRequest) error {
	r._seriesSkus = _seriesSkus
	r.Set("series_skus", _seriesSkus)
	return nil
}

// GetSeriesSkus SeriesSkus Getter
func (r AlibabawdkseriesskuremoveAPIRequest) GetSeriesSkus() *SeriesSkuRequest {
	return r._seriesSkus
}
