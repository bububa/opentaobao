package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-重置默认商品 API请求
alibaba.wdk.series.defaultsku.reset

系列品商品变更-重置默认商品
*/
type AlibabaWdkSeriesDefaultskuResetAPIRequest struct {
    model.Params
    // 系列品重置默认商品请求
    _seriesSku   *SeriesSkuRequest
}

// 初始化AlibabaWdkSeriesDefaultskuResetAPIRequest对象
func NewAlibabaWdkSeriesDefaultskuResetRequest() *AlibabaWdkSeriesDefaultskuResetAPIRequest{
    return &AlibabaWdkSeriesDefaultskuResetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesDefaultskuResetAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.defaultsku.reset"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesDefaultskuResetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SeriesSku Setter
// 系列品重置默认商品请求
func (r *AlibabaWdkSeriesDefaultskuResetAPIRequest) SetSeriesSku(_seriesSku *SeriesSkuRequest) error {
    r._seriesSku = _seriesSku
    r.Set("series_sku", _seriesSku)
    return nil
}

// SeriesSku Getter
func (r AlibabaWdkSeriesDefaultskuResetAPIRequest) GetSeriesSku() *SeriesSkuRequest {
    return r._seriesSku
}
