package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkseriesdefaultskuresetAPIRequest 系列品商品变更-重置默认商品 API请求
// alibaba.wdk.series.defaultsku.reset
//
// 系列品商品变更-重置默认商品
type AlibabawdkseriesdefaultskuresetAPIRequest struct {
	model.Params
	// 系列品重置默认商品请求
	_seriesSku *SeriesSkuRequest
}

// NewAlibabawdkseriesdefaultskuresetRequest 初始化AlibabawdkseriesdefaultskuresetAPIRequest对象
func NewAlibabawdkseriesdefaultskuresetRequest() *AlibabawdkseriesdefaultskuresetAPIRequest {
	return &AlibabawdkseriesdefaultskuresetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkseriesdefaultskuresetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.defaultsku.reset"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkseriesdefaultskuresetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkseriesdefaultskuresetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeriesSku is SeriesSku Setter
// 系列品重置默认商品请求
func (r *AlibabawdkseriesdefaultskuresetAPIRequest) SetSeriesSku(_seriesSku *SeriesSkuRequest) error {
	r._seriesSku = _seriesSku
	r.Set("series_sku", _seriesSku)
	return nil
}

// GetSeriesSku SeriesSku Getter
func (r AlibabawdkseriesdefaultskuresetAPIRequest) GetSeriesSku() *SeriesSkuRequest {
	return r._seriesSku
}
