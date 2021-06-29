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
type AlibabaWdkSeriesCreateRequest struct {
    model.Params
    // 系列品创建系列请求
    _series   *SkuSeriesCreateRequest
}

// 初始化AlibabaWdkSeriesCreateRequest对象
func NewAlibabaWdkSeriesCreateRequest() *AlibabaWdkSeriesCreateRequest{
    return &AlibabaWdkSeriesCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Series Setter
// 系列品创建系列请求
func (r *AlibabaWdkSeriesCreateRequest) SetSeries(_series *SkuSeriesCreateRequest) error {
    r._series = _series
    r.Set("series", _series)
    return nil
}

// Series Getter
func (r AlibabaWdkSeriesCreateRequest) GetSeries() *SkuSeriesCreateRequest {
    return r._series
}
