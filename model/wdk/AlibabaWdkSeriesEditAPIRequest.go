package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSeriesEditAPIRequest) Reset() {
	r._series = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesEditAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSeriesEditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeries is Series Setter
// 商品系列修改请求
func (r *AlibabaWdkSeriesEditAPIRequest) SetSeries(_series *SkuSeriesEditRequest) error {
	r._series = _series
	r.Set("series", _series)
	return nil
}

// GetSeries Series Getter
func (r AlibabaWdkSeriesEditAPIRequest) GetSeries() *SkuSeriesEditRequest {
	return r._series
}

var poolAlibabaWdkSeriesEditAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSeriesEditRequest()
	},
}

// GetAlibabaWdkSeriesEditRequest 从 sync.Pool 获取 AlibabaWdkSeriesEditAPIRequest
func GetAlibabaWdkSeriesEditAPIRequest() *AlibabaWdkSeriesEditAPIRequest {
	return poolAlibabaWdkSeriesEditAPIRequest.Get().(*AlibabaWdkSeriesEditAPIRequest)
}

// ReleaseAlibabaWdkSeriesEditAPIRequest 将 AlibabaWdkSeriesEditAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSeriesEditAPIRequest(v *AlibabaWdkSeriesEditAPIRequest) {
	v.Reset()
	poolAlibabaWdkSeriesEditAPIRequest.Put(v)
}
