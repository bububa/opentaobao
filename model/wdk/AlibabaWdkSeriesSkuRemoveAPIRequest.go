package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSeriesSkuRemoveAPIRequest) Reset() {
	r._seriesSkus = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesSkuRemoveAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.sku.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesSkuRemoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSeriesSkuRemoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeriesSkus is SeriesSkus Setter
// 系列品移除商品请求
func (r *AlibabaWdkSeriesSkuRemoveAPIRequest) SetSeriesSkus(_seriesSkus *SeriesSkuRequest) error {
	r._seriesSkus = _seriesSkus
	r.Set("series_skus", _seriesSkus)
	return nil
}

// GetSeriesSkus SeriesSkus Getter
func (r AlibabaWdkSeriesSkuRemoveAPIRequest) GetSeriesSkus() *SeriesSkuRequest {
	return r._seriesSkus
}

var poolAlibabaWdkSeriesSkuRemoveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSeriesSkuRemoveRequest()
	},
}

// GetAlibabaWdkSeriesSkuRemoveRequest 从 sync.Pool 获取 AlibabaWdkSeriesSkuRemoveAPIRequest
func GetAlibabaWdkSeriesSkuRemoveAPIRequest() *AlibabaWdkSeriesSkuRemoveAPIRequest {
	return poolAlibabaWdkSeriesSkuRemoveAPIRequest.Get().(*AlibabaWdkSeriesSkuRemoveAPIRequest)
}

// ReleaseAlibabaWdkSeriesSkuRemoveAPIRequest 将 AlibabaWdkSeriesSkuRemoveAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSeriesSkuRemoveAPIRequest(v *AlibabaWdkSeriesSkuRemoveAPIRequest) {
	v.Reset()
	poolAlibabaWdkSeriesSkuRemoveAPIRequest.Put(v)
}
