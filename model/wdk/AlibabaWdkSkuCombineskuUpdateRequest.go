package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品更新接口 APIRequest
alibaba.wdk.sku.combinesku.update

组合商品修改接口
*/
type AlibabaWdkSkuCombineskuUpdateRequest struct {
    model.Params

    // 请求参数
    paramList   []SkuDo 

}

func NewAlibabaWdkSkuCombineskuUpdateRequest() *AlibabaWdkSkuCombineskuUpdateRequest{
    return &AlibabaWdkSkuCombineskuUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuCombineskuUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.combinesku.update"
}

func (r AlibabaWdkSkuCombineskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuCombineskuUpdateRequest) SetParamList(paramList []SkuDo) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r AlibabaWdkSkuCombineskuUpdateRequest) GetParamList() []SkuDo {
    return r.paramList
}

