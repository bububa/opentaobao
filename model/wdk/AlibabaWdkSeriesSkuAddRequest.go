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
type AlibabaWdkSeriesSkuAddAPIRequest struct {
    model.Params
    // 系列品添加商品请求
    _seriesSkus   *SeriesSkuRequest
}

// 初始化AlibabaWdkSeriesSkuAddAPIRequest对象
func NewAlibabaWdkSeriesSkuAddRequest() *AlibabaWdkSeriesSkuAddAPIRequest{
    return &AlibabaWdkSeriesSkuAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesSkuAddAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.sku.add"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesSkuAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SeriesSkus Setter
// 系列品添加商品请求
func (r *AlibabaWdkSeriesSkuAddAPIRequest) SetSeriesSkus(_seriesSkus *SeriesSkuRequest) error {
    r._seriesSkus = _seriesSkus
    r.Set("series_skus", _seriesSkus)
    return nil
}

// SeriesSkus Getter
func (r AlibabaWdkSeriesSkuAddAPIRequest) GetSeriesSkus() *SeriesSkuRequest {
    return r._seriesSkus
}
