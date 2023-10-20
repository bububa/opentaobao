package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSeriesCreateAPIRequest 系列品变更-新增系列 API请求
// alibaba.wdk.series.create
//
// 系列品变更-新增系列
type AlibabaWdkSeriesCreateAPIRequest struct {
	model.Params
	// 系列品创建系列请求
	_series *SkuSeriesCreateRequest
}

// NewAlibabaWdkSeriesCreateRequest 初始化AlibabaWdkSeriesCreateAPIRequest对象
func NewAlibabaWdkSeriesCreateRequest() *AlibabaWdkSeriesCreateAPIRequest {
	return &AlibabaWdkSeriesCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSeriesCreateAPIRequest) Reset() {
	r._series = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.series.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSeriesCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSeries is Series Setter
// 系列品创建系列请求
func (r *AlibabaWdkSeriesCreateAPIRequest) SetSeries(_series *SkuSeriesCreateRequest) error {
	r._series = _series
	r.Set("series", _series)
	return nil
}

// GetSeries Series Getter
func (r AlibabaWdkSeriesCreateAPIRequest) GetSeries() *SkuSeriesCreateRequest {
	return r._series
}

var poolAlibabaWdkSeriesCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSeriesCreateRequest()
	},
}

// GetAlibabaWdkSeriesCreateRequest 从 sync.Pool 获取 AlibabaWdkSeriesCreateAPIRequest
func GetAlibabaWdkSeriesCreateAPIRequest() *AlibabaWdkSeriesCreateAPIRequest {
	return poolAlibabaWdkSeriesCreateAPIRequest.Get().(*AlibabaWdkSeriesCreateAPIRequest)
}

// ReleaseAlibabaWdkSeriesCreateAPIRequest 将 AlibabaWdkSeriesCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSeriesCreateAPIRequest(v *AlibabaWdkSeriesCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkSeriesCreateAPIRequest.Put(v)
}
