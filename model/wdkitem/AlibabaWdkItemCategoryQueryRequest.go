package wdkitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
类目查询接口 API请求
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

// 初始化AlibabaWdkItemCategoryQueryRequest对象
func NewAlibabaWdkItemCategoryQueryRequest() *AlibabaWdkItemCategoryQueryRequest{
    return &AlibabaWdkItemCategoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemCategoryQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.category.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemCategoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keyword Setter
// 查询关键词，不填查全部
func (r *AlibabaWdkItemCategoryQueryRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

// Keyword Getter
func (r AlibabaWdkItemCategoryQueryRequest) GetKeyword() string {
    return r.keyword
}
// RootCategoryCode Setter
// 从哪个类目开始查，不填从根类目开始查
func (r *AlibabaWdkItemCategoryQueryRequest) SetRootCategoryCode(rootCategoryCode string) error {
    r.rootCategoryCode = rootCategoryCode
    r.Set("root_category_code", rootCategoryCode)
    return nil
}

// RootCategoryCode Getter
func (r AlibabaWdkItemCategoryQueryRequest) GetRootCategoryCode() string {
    return r.rootCategoryCode
}
