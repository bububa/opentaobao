package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerproductsearchAPIRequest 查询商品api API请求
// alibaba.alihealth.tracecodeseller.product.search
//
// 查询商品api
type AlibabaalihealthtracecodesellerproductsearchAPIRequest struct {
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

// NewAlibabaalihealthtracecodesellerproductsearchRequest 初始化AlibabaalihealthtracecodesellerproductsearchAPIRequest对象
func NewAlibabaalihealthtracecodesellerproductsearchRequest() *AlibabaalihealthtracecodesellerproductsearchAPIRequest {
	return &AlibabaalihealthtracecodesellerproductsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesellerproductsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.product.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesellerproductsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesellerproductsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkeyCode is SkeyCode Setter
// 身份认证
func (r *AlibabaalihealthtracecodesellerproductsearchAPIRequest) SetSkeyCode(_skeyCode string) error {
	r._skeyCode = _skeyCode
	r.Set("skey_code", _skeyCode)
	return nil
}

// GetSkeyCode SkeyCode Getter
func (r AlibabaalihealthtracecodesellerproductsearchAPIRequest) GetSkeyCode() string {
	return r._skeyCode
}

// SetEntInfoId is EntInfoId Setter
// 商家id
func (r *AlibabaalihealthtracecodesellerproductsearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaalihealthtracecodesellerproductsearchAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// SetPage is Page Setter
// 页数
func (r *AlibabaalihealthtracecodesellerproductsearchAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaalihealthtracecodesellerproductsearchAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *AlibabaalihealthtracecodesellerproductsearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthtracecodesellerproductsearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
