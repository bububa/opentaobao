package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新建商品草稿 APIRequest
alibaba.wdk.merchant.item.createdraft

新建商品草稿erp接口
*/
type AlibabaWdkMerchantItemCreatedraftRequest struct {
    model.Params

    // 商品信息json
    params   string 

}

func NewAlibabaWdkMerchantItemCreatedraftRequest() *AlibabaWdkMerchantItemCreatedraftRequest{
    return &AlibabaWdkMerchantItemCreatedraftRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMerchantItemCreatedraftRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.item.createdraft"
}

func (r AlibabaWdkMerchantItemCreatedraftRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMerchantItemCreatedraftRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r AlibabaWdkMerchantItemCreatedraftRequest) GetParams() string {
    return r.params
}

