package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品-商品排序 APIRequest
alibaba.wdk.series.sort

系列品商品变更-商品排序
*/
type AlibabaWdkSeriesSortRequest struct {
    model.Params

    // 自定义排序请求
    sort   *SeriesSortRequest 

}

func NewAlibabaWdkSeriesSortRequest() *AlibabaWdkSeriesSortRequest{
    return &AlibabaWdkSeriesSortRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSeriesSortRequest) GetApiMethodName() string {
    return "alibaba.wdk.series.sort"
}

func (r AlibabaWdkSeriesSortRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSeriesSortRequest) SetSort(sort *SeriesSortRequest) error {
    r.sort = sort
    r.Set("sort", sort)
    return nil
}

func (r AlibabaWdkSeriesSortRequest) GetSort() *SeriesSortRequest {
    return r.sort
}

