package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesDefaultskuResetAPIRequest 系列品商品变更-重置默认商品 API请求
// alibaba.wdk.series.defaultsku.reset
//
// 系列品商品变更-重置默认商品
type AlibabaWdkSeriesDefaultskuResetAPIRequest struct {
	model.Params
	// 系列品重置默认商品请求
	_seriesSku *SeriesSkuRequest
}

// NewAlibabaWdkSeriesDefaultskuResetRequest 初始化AlibabaWdkSeriesDefaultskuResetAPIRequest对象
func NewAlibabaWdkSeriesDefaultskuResetRequest() *AlibabaWdkSeriesDefaultskuResetAPIRequest {
	return &AlibabaWdkSeriesDefaultskuResetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSeriesDefaultskuResetAPIRequest) Reset() {
	r._seriesSku = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesDefaultskuResetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.defaultsku.reset"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesDefaultskuResetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSeriesDefaultskuResetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeriesSku is SeriesSku Setter
// 系列品重置默认商品请求
func (r *AlibabaWdkSeriesDefaultskuResetAPIRequest) SetSeriesSku(_seriesSku *SeriesSkuRequest) error {
	r._seriesSku = _seriesSku
	r.Set("series_sku", _seriesSku)
	return nil
}

// GetSeriesSku SeriesSku Getter
func (r AlibabaWdkSeriesDefaultskuResetAPIRequest) GetSeriesSku() *SeriesSkuRequest {
	return r._seriesSku
}

var poolAlibabaWdkSeriesDefaultskuResetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSeriesDefaultskuResetRequest()
	},
}

// GetAlibabaWdkSeriesDefaultskuResetRequest 从 sync.Pool 获取 AlibabaWdkSeriesDefaultskuResetAPIRequest
func GetAlibabaWdkSeriesDefaultskuResetAPIRequest() *AlibabaWdkSeriesDefaultskuResetAPIRequest {
	return poolAlibabaWdkSeriesDefaultskuResetAPIRequest.Get().(*AlibabaWdkSeriesDefaultskuResetAPIRequest)
}

// ReleaseAlibabaWdkSeriesDefaultskuResetAPIRequest 将 AlibabaWdkSeriesDefaultskuResetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSeriesDefaultskuResetAPIRequest(v *AlibabaWdkSeriesDefaultskuResetAPIRequest) {
	v.Reset()
	poolAlibabaWdkSeriesDefaultskuResetAPIRequest.Put(v)
}
