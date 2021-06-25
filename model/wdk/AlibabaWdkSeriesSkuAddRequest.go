package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-添加商品 APIRequest
alibaba.wdk.series.sku.add

系列品商品变更-添加商品
*/
type AlibabaWdkSeriesSkuAddRequest struct {
    model.Params

    // 系列品添加商品请求
    seriesSkus   *SeriesSkuRequest 

}

func NewAlibabaWdkSeriesSkuAddRequest() *AlibabaWdkSeriesSkuAddRequest{
    return &AlibabaWdkSeriesSkuAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSeriesSkuAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.sku.add"
}

func (r AlibabaWdkSeriesSkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSeriesSkuAddRequest) SetSeriesSkus(seriesSkus *SeriesSkuRequest) error {
    r.seriesSkus = seriesSkus
    r.Set("series_skus", seriesSkus)
    return nil
}

func (r AlibabaWdkSeriesSkuAddRequest) GetSeriesSkus() *SeriesSkuRequest {
    return r.seriesSkus
}

