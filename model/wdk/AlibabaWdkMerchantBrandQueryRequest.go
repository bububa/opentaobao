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
    _keyword   string
    // 可不填，默认0
    _offset   int64
    // 可不填，默认2000
    _pageSize   int64
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
func (r *AlibabaWdkMerchantBrandQueryRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r AlibabaWdkMerchantBrandQueryRequest) GetKeyword() string {
    return r._keyword
}
// Offset Setter
// 可不填，默认0
func (r *AlibabaWdkMerchantBrandQueryRequest) SetOffset(_offset int64) error {
    r._offset = _offset
    r.Set("offset", _offset)
    return nil
}

// Offset Getter
func (r AlibabaWdkMerchantBrandQueryRequest) GetOffset() int64 {
    return r._offset
}
// PageSize Setter
// 可不填，默认2000
func (r *AlibabaWdkMerchantBrandQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaWdkMerchantBrandQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
