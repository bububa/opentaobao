package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌查询接口 API请求
alibaba.wdk.merchant.brand.query

三江erp对接时，提供品牌查询的接口
*/
type AlibabaWdkMerchantBrandQueryRequest struct {
    model.Params
    // 关键词，不填就查全部
    keyword   string
    // 可不填，默认0
    offset   int64
    // 可不填，默认2000
    pageSize   int64
}

// 初始化AlibabaWdkMerchantBrandQueryRequest对象
func NewAlibabaWdkMerchantBrandQueryRequest() *AlibabaWdkMerchantBrandQueryRequest{
    return &AlibabaWdkMerchantBrandQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMerchantBrandQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.merchant.brand.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMerchantBrandQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keyword Setter
// 关键词，不填就查全部
func (r *AlibabaWdkMerchantBrandQueryRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

// Keyword Getter
func (r AlibabaWdkMerchantBrandQueryRequest) GetKeyword() string {
    return r.keyword
}
// Offset Setter
// 可不填，默认0
func (r *AlibabaWdkMerchantBrandQueryRequest) SetOffset(offset int64) error {
    r.offset = offset
    r.Set("offset", offset)
    return nil
}

// Offset Getter
func (r AlibabaWdkMerchantBrandQueryRequest) GetOffset() int64 {
    return r.offset
}
// PageSize Setter
// 可不填，默认2000
func (r *AlibabaWdkMerchantBrandQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaWdkMerchantBrandQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
