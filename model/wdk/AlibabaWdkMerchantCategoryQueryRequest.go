package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三江erp对接类目查询接口 APIRequest
alibaba.wdk.merchant.category.query

三江erp对接类目查询接口
*/
type AlibabaWdkMerchantCategoryQueryRequest struct {
    model.Params

    // 搜索关键词，可不填就查全部
    keyword   string 

    // 类目起点，可不填从根目录开始查
    rootCategoryCode   string 

}

func NewAlibabaWdkMerchantCategoryQueryRequest() *AlibabaWdkMerchantCategoryQueryRequest{
    return &AlibabaWdkMerchantCategoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMerchantCategoryQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.category.query"
}

func (r AlibabaWdkMerchantCategoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMerchantCategoryQueryRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r AlibabaWdkMerchantCategoryQueryRequest) GetKeyword() string {
    return r.keyword
}

func (r *AlibabaWdkMerchantCategoryQueryRequest) SetRootCategoryCode(rootCategoryCode string) error {
    r.rootCategoryCode = rootCategoryCode
    r.Set("root_category_code", rootCategoryCode)
    return nil
}

func (r AlibabaWdkMerchantCategoryQueryRequest) GetRootCategoryCode() string {
    return r.rootCategoryCode
}

