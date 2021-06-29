package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品-商品排序 API请求
alibaba.wdk.series.sort

系列品商品变更-商品排序
*/
type AlibabaWdkSeriesSortRequest struct {
    model.Params
    // 自定义排序请求
    sort   *SeriesSortRequest
}

// 初始化AlibabaWdkSeriesSortRequest对象
func NewAlibabaWdkSeriesSortRequest() *AlibabaWdkSeriesSortRequest{
    return &AlibabaWdkSeriesSortRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkSeriesSortRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.sort"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkSeriesSortRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Sort Setter
// 自定义排序请求
func (r *AlibabaWdkSeriesSortRequest) SetSort(sort *SeriesSortRequest) error {
    r.sort = sort
    r.Set("sort", sort)
    return nil
}

// Sort Getter
func (r AlibabaWdkSeriesSortRequest) GetSort() *SeriesSortRequest {
    return r.sort
}
