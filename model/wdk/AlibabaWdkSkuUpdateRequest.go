package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品 APIRequest
alibaba.wdk.sku.update

开放商品更新接口
*/
type AlibabaWdkSkuUpdateRequest struct {
    model.Params

    // 参数列表
    paramList   []SkuDo 

}

func NewAlibabaWdkSkuUpdateRequest() *AlibabaWdkSkuUpdateRequest{
    return &AlibabaWdkSkuUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.update"
}

func (r AlibabaWdkSkuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuUpdateRequest) SetParamList(paramList []SkuDo) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r AlibabaWdkSkuUpdateRequest) GetParamList() []SkuDo {
    return r.paramList
}

