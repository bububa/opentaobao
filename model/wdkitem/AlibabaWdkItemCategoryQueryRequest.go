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
type AlibabaWdkItemCategoryQueryAPIRequest struct {
    model.Params
    // 查询关键词，不填查全部
    _keyword   string
    // 从哪个类目开始查，不填从根类目开始查
    _rootCategoryCode   string
}

// 初始化AlibabaWdkItemCategoryQueryAPIRequest对象
func NewAlibabaWdkItemCategoryQueryRequest() *AlibabaWdkItemCategoryQueryAPIRequest{
    return &AlibabaWdkItemCategoryQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemCategoryQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.item.category.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemCategoryQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keyword Setter
// 查询关键词，不填查全部
func (r *AlibabaWdkItemCategoryQueryAPIRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r AlibabaWdkItemCategoryQueryAPIRequest) GetKeyword() string {
    return r._keyword
}
// RootCategoryCode Setter
// 从哪个类目开始查，不填从根类目开始查
func (r *AlibabaWdkItemCategoryQueryAPIRequest) SetRootCategoryCode(_rootCategoryCode string) error {
    r._rootCategoryCode = _rootCategoryCode
    r.Set("root_category_code", _rootCategoryCode)
    return nil
}

// RootCategoryCode Getter
func (r AlibabaWdkItemCategoryQueryAPIRequest) GetRootCategoryCode() string {
    return r._rootCategoryCode
}
