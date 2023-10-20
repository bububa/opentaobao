package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesSkuAddAPIRequest 系列品商品变更-添加商品 API请求
// alibaba.wdk.series.sku.add
//
// 系列品商品变更-添加商品
type AlibabaWdkSeriesSkuAddAPIRequest struct {
	model.Params
	// 系列品添加商品请求
	_seriesSkus *SeriesSkuRequest
}

// NewAlibabaWdkSeriesSkuAddRequest 初始化AlibabaWdkSeriesSkuAddAPIRequest对象
func NewAlibabaWdkSeriesSkuAddRequest() *AlibabaWdkSeriesSkuAddAPIRequest {
	return &AlibabaWdkSeriesSkuAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSeriesSkuAddAPIRequest) Reset() {
	r._seriesSkus = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesSkuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesSkuAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSeriesSkuAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeriesSkus is SeriesSkus Setter
// 系列品添加商品请求
func (r *AlibabaWdkSeriesSkuAddAPIRequest) SetSeriesSkus(_seriesSkus *SeriesSkuRequest) error {
	r._seriesSkus = _seriesSkus
	r.Set("series_skus", _seriesSkus)
	return nil
}

// GetSeriesSkus SeriesSkus Getter
func (r AlibabaWdkSeriesSkuAddAPIRequest) GetSeriesSkus() *SeriesSkuRequest {
	return r._seriesSkus
}

var poolAlibabaWdkSeriesSkuAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSeriesSkuAddRequest()
	},
}

// GetAlibabaWdkSeriesSkuAddRequest 从 sync.Pool 获取 AlibabaWdkSeriesSkuAddAPIRequest
func GetAlibabaWdkSeriesSkuAddAPIRequest() *AlibabaWdkSeriesSkuAddAPIRequest {
	return poolAlibabaWdkSeriesSkuAddAPIRequest.Get().(*AlibabaWdkSeriesSkuAddAPIRequest)
}

// ReleaseAlibabaWdkSeriesSkuAddAPIRequest 将 AlibabaWdkSeriesSkuAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSeriesSkuAddAPIRequest(v *AlibabaWdkSeriesSkuAddAPIRequest) {
	v.Reset()
	poolAlibabaWdkSeriesSkuAddAPIRequest.Put(v)
}
