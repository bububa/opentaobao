package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品变更-更新系列 APIRequest
alibaba.wdk.series.edit

系列品变更-更新系列
*/
type AlibabaWdkSeriesEditRequest struct {
    model.Params

    // 商品系列修改请求
    series   *SkuSeriesEditRequest 

}

func NewAlibabaWdkSeriesEditRequest() *AlibabaWdkSeriesEditRequest{
    return &AlibabaWdkSeriesEditRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSeriesEditRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.edit"
}

func (r AlibabaWdkSeriesEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSeriesEditRequest) SetSeries(series *SkuSeriesEditRequest) error {
    r.series = series
    r.Set("series", series)
    return nil
}

func (r AlibabaWdkSeriesEditRequest) GetSeries() *SkuSeriesEditRequest {
    return r.series
}

