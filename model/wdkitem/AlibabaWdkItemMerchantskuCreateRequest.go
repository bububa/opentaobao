package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品信息新建 APIRequest
alibaba.wdk.item.merchantsku.create

商家商品信息新建
*/
type AlibabaWdkItemMerchantskuCreateRequest struct {
    model.Params

    // 新建商品参数，是个json字符串
    params   string 

}

func NewAlibabaWdkItemMerchantskuCreateRequest() *AlibabaWdkItemMerchantskuCreateRequest{
    return &AlibabaWdkItemMerchantskuCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemMerchantskuCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchantsku.create"
}

func (r AlibabaWdkItemMerchantskuCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemMerchantskuCreateRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r AlibabaWdkItemMerchantskuCreateRequest) GetParams() string {
    return r.params
}

