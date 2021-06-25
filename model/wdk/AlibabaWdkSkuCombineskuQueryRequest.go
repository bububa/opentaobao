package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品查询接口 APIRequest
alibaba.wdk.sku.combinesku.query

查询组合商品接口
*/
type AlibabaWdkSkuCombineskuQueryRequest struct {
    model.Params

    // 请求参数
    param   *SkuQueryDO 

}

func NewAlibabaWdkSkuCombineskuQueryRequest() *AlibabaWdkSkuCombineskuQueryRequest{
    return &AlibabaWdkSkuCombineskuQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuCombineskuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.combinesku.query"
}

func (r AlibabaWdkSkuCombineskuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuCombineskuQueryRequest) SetParam(param *SkuQueryDO) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkSkuCombineskuQueryRequest) GetParam() *SkuQueryDO {
    return r.param
}

