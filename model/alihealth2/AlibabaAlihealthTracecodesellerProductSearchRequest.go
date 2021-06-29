package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品api API请求
alibaba.alihealth.tracecodeseller.product.search

查询商品api
*/
type AlibabaAlihealthTracecodesellerProductSearchRequest struct {
    model.Params
    // 身份认证
    skeyCode   string
    // 商家id
    entInfoId   int64
    // 页数
    page   int64
    // 每页条数
    pageSize   int64
}

// 初始化AlibabaAlihealthTracecodesellerProductSearchRequest对象
func NewAlibabaAlihealthTracecodesellerProductSearchRequest() *AlibabaAlihealthTracecodesellerProductSearchRequest{
    return &AlibabaAlihealthTracecodesellerProductSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerProductSearchRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.product.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerProductSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkeyCode Setter
// 身份认证
func (r *AlibabaAlihealthTracecodesellerProductSearchRequest) SetSkeyCode(skeyCode string) error {
    r.skeyCode = skeyCode
    r.Set("skey_code", skeyCode)
    return nil
}

// SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerProductSearchRequest) GetSkeyCode() string {
    return r.skeyCode
}
// EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerProductSearchRequest) SetEntInfoId(entInfoId int64) error {
    r.entInfoId = entInfoId
    r.Set("ent_info_id", entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerProductSearchRequest) GetEntInfoId() int64 {
    return r.entInfoId
}
// Page Setter
// 页数
func (r *AlibabaAlihealthTracecodesellerProductSearchRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthTracecodesellerProductSearchRequest) GetPage() int64 {
    return r.page
}
// PageSize Setter
// 每页条数
func (r *AlibabaAlihealthTracecodesellerProductSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthTracecodesellerProductSearchRequest) GetPageSize() int64 {
    return r.pageSize
}
