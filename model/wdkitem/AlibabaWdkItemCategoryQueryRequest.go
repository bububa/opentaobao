package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
类目查询接口 APIRequest
alibaba.wdk.item.category.query

类目查询接口
*/
type AlibabaWdkItemCategoryQueryRequest struct {
    model.Params

    // 查询关键词，不填查全部
    keyword   string 

    // 从哪个类目开始查，不填从根类目开始查
    rootCategoryCode   string 

}

func NewAlibabaWdkItemCategoryQueryRequest() *AlibabaWdkItemCategoryQueryRequest{
    return &AlibabaWdkItemCategoryQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkItemCategoryQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.category.query"
}

func (r AlibabaWdkItemCategoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkItemCategoryQueryRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r AlibabaWdkItemCategoryQueryRequest) GetKeyword() string {
    return r.keyword
}

func (r *AlibabaWdkItemCategoryQueryRequest) SetRootCategoryCode(rootCategoryCode string) error {
    r.rootCategoryCode = rootCategoryCode
    r.Set("root_category_code", rootCategoryCode)
    return nil
}

func (r AlibabaWdkItemCategoryQueryRequest) GetRootCategoryCode() string {
    return r.rootCategoryCode
}

