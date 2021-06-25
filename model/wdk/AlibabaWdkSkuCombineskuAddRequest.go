package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品新增接口 APIRequest
alibaba.wdk.sku.combinesku.add

组合商品新增接口
*/
type AlibabaWdkSkuCombineskuAddRequest struct {
    model.Params

    // 请求参数
    paramList   []SkuDo 

}

func NewAlibabaWdkSkuCombineskuAddRequest() *AlibabaWdkSkuCombineskuAddRequest{
    return &AlibabaWdkSkuCombineskuAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuCombineskuAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.combinesku.add"
}

func (r AlibabaWdkSkuCombineskuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuCombineskuAddRequest) SetParamList(paramList []SkuDo) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r AlibabaWdkSkuCombineskuAddRequest) GetParamList() []SkuDo {
    return r.paramList
}

