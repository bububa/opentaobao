package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品变更-新增系列 API请求
alibaba.wdk.series.create

系列品变更-新增系列
*/
type AlibabaWdkSeriesCreateAPIRequest struct {
    model.Params
    // 系列品创建系列请求
    _series   *SkuSeriesCreateRequest
}

// 初始化AlibabaWdkSeriesCreateAPIRequest对象
func NewAlibabaWdkSeriesCreateRequest() *AlibabaWdkSeriesCreateAPIRequest{
    return &AlibabaWdkSeriesCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Series Setter
// 系列品创建系列请求
func (r *AlibabaWdkSeriesCreateAPIRequest) SetSeries(_series *SkuSeriesCreateRequest) error {
    r._series = _series
    r.Set("series", _series)
    return nil
}

// Series Getter
func (r AlibabaWdkSeriesCreateAPIRequest) GetSeries() *SkuSeriesCreateRequest {
    return r._series
}
