package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-移除商品 APIRequest
alibaba.wdk.series.sku.remove

系列品商品变更-移除商品
*/
type AlibabaWdkSeriesSkuRemoveRequest struct {
    model.Params

    // 系列品移除商品请求
    seriesSkus   *SeriesSkuRequest 

}

func NewAlibabaWdkSeriesSkuRemoveRequest() *AlibabaWdkSeriesSkuRemoveRequest{
    return &AlibabaWdkSeriesSkuRemoveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSeriesSkuRemoveRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.sku.remove"
}

func (r AlibabaWdkSeriesSkuRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSeriesSkuRemoveRequest) SetSeriesSkus(seriesSkus *SeriesSkuRequest) error {
    r.seriesSkus = seriesSkus
    r.Set("series_skus", seriesSkus)
    return nil
}

func (r AlibabaWdkSeriesSkuRemoveRequest) GetSeriesSkus() *SeriesSkuRequest {
    return r.seriesSkus
}

