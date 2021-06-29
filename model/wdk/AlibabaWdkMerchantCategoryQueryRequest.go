package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三江erp对接类目查询接口 API请求
alibaba.wdk.merchant.category.query

三江erp对接类目查询接口
*/
type AlibabaWdkMerchantCategoryQueryRequest struct {
    model.Params
    // 搜索关键词，可不填就查全部
    _keyword   string
    // 类目起点，可不填从根目录开始查
    _rootCategoryCode   string
}

// 初始化AlibabaWdkMerchantCategoryQueryRequest对象
func NewAlibabaWdkMerchantCategoryQueryRequest() *AlibabaWdkMerchantCategoryQueryRequest{
    return &AlibabaWdkMerchantCategoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantCategoryQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.category.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantCategoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keyword Setter
// 搜索关键词，可不填就查全部
func (r *AlibabaWdkMerchantCategoryQueryRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r AlibabaWdkMerchantCategoryQueryRequest) GetKeyword() string {
    return r._keyword
}
// RootCategoryCode Setter
// 类目起点，可不填从根目录开始查
func (r *AlibabaWdkMerchantCategoryQueryRequest) SetRootCategoryCode(_rootCategoryCode string) error {
    r._rootCategoryCode = _rootCategoryCode
    r.Set("root_category_code", _rootCategoryCode)
    return nil
}

// RootCategoryCode Getter
func (r AlibabaWdkMerchantCategoryQueryRequest) GetRootCategoryCode() string {
    return r._rootCategoryCode
}
