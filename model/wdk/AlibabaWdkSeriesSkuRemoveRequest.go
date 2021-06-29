package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-移除商品 API请求
alibaba.wdk.series.sku.remove

系列品商品变更-移除商品
*/
type AlibabaWdkSeriesSkuRemoveRequest struct {
    model.Params
    // 系列品移除商品请求
    seriesSkus   *SeriesSkuRequest
}

// 初始化AlibabaWdkSeriesSkuRemoveRequest对象
func NewAlibabaWdkSeriesSkuRemoveRequest() *AlibabaWdkSeriesSkuRemoveRequest{
    return &AlibabaWdkSeriesSkuRemoveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesSkuRemoveRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.sku.remove"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesSkuRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SeriesSkus Setter
// 系列品移除商品请求
func (r *AlibabaWdkSeriesSkuRemoveRequest) SetSeriesSkus(seriesSkus *SeriesSkuRequest) error {
    r.seriesSkus = seriesSkus
    r.Set("series_skus", seriesSkus)
    return nil
}

// SeriesSkus Getter
func (r AlibabaWdkSeriesSkuRemoveRequest) GetSeriesSkus() *SeriesSkuRequest {
    return r.seriesSkus
}
