package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-添加商品 API请求
alibaba.wdk.series.sku.add

系列品商品变更-添加商品
*/
type AlibabaWdkSeriesSkuAddRequest struct {
    model.Params
    // 系列品添加商品请求
    seriesSkus   *SeriesSkuRequest
}

// 初始化AlibabaWdkSeriesSkuAddRequest对象
func NewAlibabaWdkSeriesSkuAddRequest() *AlibabaWdkSeriesSkuAddRequest{
    return &AlibabaWdkSeriesSkuAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesSkuAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.sku.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesSkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SeriesSkus Setter
// 系列品添加商品请求
func (r *AlibabaWdkSeriesSkuAddRequest) SetSeriesSkus(seriesSkus *SeriesSkuRequest) error {
    r.seriesSkus = seriesSkus
    r.Set("series_skus", seriesSkus)
    return nil
}

// SeriesSkus Getter
func (r AlibabaWdkSeriesSkuAddRequest) GetSeriesSkus() *SeriesSkuRequest {
    return r.seriesSkus
}
