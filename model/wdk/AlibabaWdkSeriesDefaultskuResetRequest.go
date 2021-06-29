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
type AlibabaWdkSeriesDefaultskuResetRequest struct {
    model.Params
    // 系列品重置默认商品请求
    _seriesSku   *SeriesSkuRequest
}

// 初始化AlibabaWdkSeriesDefaultskuResetRequest对象
func NewAlibabaWdkSeriesDefaultskuResetRequest() *AlibabaWdkSeriesDefaultskuResetRequest{
    return &AlibabaWdkSeriesDefaultskuResetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesDefaultskuResetRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.defaultsku.reset"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesDefaultskuResetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SeriesSku Setter
// 系列品重置默认商品请求
func (r *AlibabaWdkSeriesDefaultskuResetRequest) SetSeriesSku(_seriesSku *SeriesSkuRequest) error {
    r._seriesSku = _seriesSku
    r.Set("series_sku", _seriesSku)
    return nil
}

// SeriesSku Getter
func (r AlibabaWdkSeriesDefaultskuResetRequest) GetSeriesSku() *SeriesSkuRequest {
    return r._seriesSku
}
