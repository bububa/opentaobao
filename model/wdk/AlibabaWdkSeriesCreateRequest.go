package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品变更-新增系列 APIRequest
alibaba.wdk.series.create

系列品变更-新增系列
*/
type AlibabaWdkSeriesCreateRequest struct {
    model.Params

    // 系列品创建系列请求
    series   *SkuSeriesCreateRequest 

}

func NewAlibabaWdkSeriesCreateRequest() *AlibabaWdkSeriesCreateRequest{
    return &AlibabaWdkSeriesCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSeriesCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.create"
}

func (r AlibabaWdkSeriesCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSeriesCreateRequest) SetSeries(series *SkuSeriesCreateRequest) error {
    r.series = series
    r.Set("series", series)
    return nil
}

func (r AlibabaWdkSeriesCreateRequest) GetSeries() *SkuSeriesCreateRequest {
    return r.series
}

