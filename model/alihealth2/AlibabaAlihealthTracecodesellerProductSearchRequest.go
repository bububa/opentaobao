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
    _skeyCode   string
    // 商家id
    _entInfoId   int64
    // 页数
    _page   int64
    // 每页条数
    _pageSize   int64
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
func (r *AlibabaAlihealthTracecodesellerProductSearchRequest) SetSkeyCode(_skeyCode string) error {
    r._skeyCode = _skeyCode
    r.Set("skey_code", _skeyCode)
    return nil
}

// SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerProductSearchRequest) GetSkeyCode() string {
    return r._skeyCode
}
// EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerProductSearchRequest) SetEntInfoId(_entInfoId int64) error {
    r._entInfoId = _entInfoId
    r.Set("ent_info_id", _entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerProductSearchRequest) GetEntInfoId() int64 {
    return r._entInfoId
}
// Page Setter
// 页数
func (r *AlibabaAlihealthTracecodesellerProductSearchRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthTracecodesellerProductSearchRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 每页条数
func (r *AlibabaAlihealthTracecodesellerProductSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthTracecodesellerProductSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
