package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-重置默认商品 APIRequest
alibaba.wdk.series.defaultsku.reset

系列品商品变更-重置默认商品
*/
type AlibabaWdkSeriesDefaultskuResetRequest struct {
    model.Params

    // 系列品重置默认商品请求
    seriesSku   *SeriesSkuRequest 

}

func NewAlibabaWdkSeriesDefaultskuResetRequest() *AlibabaWdkSeriesDefaultskuResetRequest{
    return &AlibabaWdkSeriesDefaultskuResetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSeriesDefaultskuResetRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.defaultsku.reset"
}

func (r AlibabaWdkSeriesDefaultskuResetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSeriesDefaultskuResetRequest) SetSeriesSku(seriesSku *SeriesSkuRequest) error {
    r.seriesSku = seriesSku
    r.Set("series_sku", seriesSku)
    return nil
}

func (r AlibabaWdkSeriesDefaultskuResetRequest) GetSeriesSku() *SeriesSkuRequest {
    return r.seriesSku
}

