package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerProductSearchAPIRequest
查询商品api API请求
alibaba.alihealth.tracecodeseller.product.search

查询商品api */
type AlibabaAlihealthTracecodesellerProductSearchAPIRequest struct {
	model.Params
	// 身份认证
	_skeyCode string
	// 商家id
	_entInfoId int64
	// 页数
	_page int64
	// 每页条数
	_pageSize int64
}

// NewAlibabaAlihealthTracecodesellerProductSearchRequest 初始化AlibabaAlihealthTracecodesellerProductSearchAPIRequest对象
func NewAlibabaAlihealthTracecodesellerProductSearchRequest() *AlibabaAlihealthTracecodesellerProductSearchAPIRequest {
	return &AlibabaAlihealthTracecodesellerProductSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerProductSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.product.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerProductSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SkeyCode Setter
// 身份认证
func (r *AlibabaAlihealthTracecodesellerProductSearchAPIRequest) SetSkeyCode(_skeyCode string) error {
	r._skeyCode = _skeyCode
	r.Set("skey_code", _skeyCode)
	return nil
}

// Get SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerProductSearchAPIRequest) GetSkeyCode() string {
	return r._skeyCode
}

// Set is EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerProductSearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// Get EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerProductSearchAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// Set is Page Setter
// 页数
func (r *AlibabaAlihealthTracecodesellerProductSearchAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaAlihealthTracecodesellerProductSearchAPIRequest) GetPage() int64 {
	return r._page
}

// Set is PageSize Setter
// 每页条数
func (r *AlibabaAlihealthTracecodesellerProductSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAlihealthTracecodesellerProductSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
