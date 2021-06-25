package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增商品 APIRequest
alibaba.wdk.sku.add

创建RT门店商品或DC商品
*/
type AlibabaWdkSkuAddRequest struct {
    model.Params

    // 商品列表
    paramList   []SkuDo 

}

func NewAlibabaWdkSkuAddRequest() *AlibabaWdkSkuAddRequest{
    return &AlibabaWdkSkuAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuAddRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.add"
}

func (r AlibabaWdkSkuAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuAddRequest) SetParamList(paramList []SkuDo) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r AlibabaWdkSkuAddRequest) GetParamList() []SkuDo {
    return r.paramList
}

