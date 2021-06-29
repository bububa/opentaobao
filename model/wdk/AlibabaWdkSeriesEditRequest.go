package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品变更-更新系列 API请求
alibaba.wdk.series.edit

系列品变更-更新系列
*/
type AlibabaWdkSeriesEditRequest struct {
    model.Params
    // 商品系列修改请求
    series   *SkuSeriesEditRequest
}

// 初始化AlibabaWdkSeriesEditRequest对象
func NewAlibabaWdkSeriesEditRequest() *AlibabaWdkSeriesEditRequest{
    return &AlibabaWdkSeriesEditRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesEditRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.edit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesEditRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Series Setter
// 商品系列修改请求
func (r *AlibabaWdkSeriesEditRequest) SetSeries(series *SkuSeriesEditRequest) error {
    r.series = series
    r.Set("series", series)
    return nil
}

// Series Getter
func (r AlibabaWdkSeriesEditRequest) GetSeries() *SkuSeriesEditRequest {
    return r.series
}
