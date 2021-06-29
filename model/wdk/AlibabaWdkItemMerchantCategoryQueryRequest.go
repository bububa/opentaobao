package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品的商家叶子类目 API请求
alibaba.wdk.item.merchant.category.query

查询商品的商家叶子类目
*/
type AlibabaWdkItemMerchantCategoryQueryRequest struct {
    model.Params
    // 请求
    queryRequest   *WdkOpenSkuMerchantCatServiceQueryRequest
}

// 初始化AlibabaWdkItemMerchantCategoryQueryRequest对象
func NewAlibabaWdkItemMerchantCategoryQueryRequest() *AlibabaWdkItemMerchantCategoryQueryRequest{
    return &AlibabaWdkItemMerchantCategoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMerchantCategoryQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.merchant.category.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMerchantCategoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryRequest Setter
// 请求
func (r *AlibabaWdkItemMerchantCategoryQueryRequest) SetQueryRequest(queryRequest *WdkOpenSkuMerchantCatServiceQueryRequest) error {
    r.queryRequest = queryRequest
    r.Set("query_request", queryRequest)
    return nil
}

// QueryRequest Getter
func (r AlibabaWdkItemMerchantCategoryQueryRequest) GetQueryRequest() *WdkOpenSkuMerchantCatServiceQueryRequest {
    return r.queryRequest
}
