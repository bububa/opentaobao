package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品 APIRequest
alibaba.wdk.sku.query

查询商品
*/
type AlibabaWdkSkuQueryRequest struct {
    model.Params

    // 入参
    param   *SkuQueryDo 

}

func NewAlibabaWdkSkuQueryRequest() *AlibabaWdkSkuQueryRequest{
    return &AlibabaWdkSkuQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.query"
}

func (r AlibabaWdkSkuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuQueryRequest) SetParam(param *SkuQueryDo) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkSkuQueryRequest) GetParam() *SkuQueryDo {
    return r.param
}

