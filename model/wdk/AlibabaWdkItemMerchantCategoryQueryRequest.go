package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品的商家叶子类目 APIRequest
alibaba.wdk.item.merchant.category.query

查询商品的商家叶子类目
*/
type AlibabaWdkItemMerchantCategoryQueryRequest struct {
    model.Params

    // 请求
    queryRequest   *WdkOpenSkuMerchantCatServiceQueryRequest 

}

func NewAlibabaWdkItemMerchantCategoryQueryRequest() *AlibabaWdkItemMerchantCategoryQueryRequest{
    return &AlibabaWdkItemMerchantCategoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemMerchantCategoryQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchant.category.query"
}

func (r AlibabaWdkItemMerchantCategoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemMerchantCategoryQueryRequest) SetQueryRequest(queryRequest *WdkOpenSkuMerchantCatServiceQueryRequest) error {
    r.queryRequest = queryRequest
    r.Set("query_request", queryRequest)
    return nil
}

func (r AlibabaWdkItemMerchantCategoryQueryRequest) GetQueryRequest() *WdkOpenSkuMerchantCatServiceQueryRequest {
    return r.queryRequest
}

