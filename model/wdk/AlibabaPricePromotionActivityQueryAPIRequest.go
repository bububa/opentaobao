package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPricePromotionActivityQueryAPIRequest 查询盒马帮档期活动详情 API请求
// alibaba.price.promotion.activity.query
//
// 查询盒马帮档期活动详情
type AlibabaPricePromotionActivityQueryAPIRequest struct {
	model.Params
	// 外部档期code
	_outerPromotionCode string
	// TOB店仓编码
	_ouCode string
	// 页码
	_page int64
	// 页码大小
	_pageSize int64
}

// NewAlibabaPricePromotionActivityQueryRequest 初始化AlibabaPricePromotionActivityQueryAPIRequest对象
func NewAlibabaPricePromotionActivityQueryRequest() *AlibabaPricePromotionActivityQueryAPIRequest {
	return &AlibabaPricePromotionActivityQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPricePromotionActivityQueryAPIRequest) Reset() {
	r._outerPromotionCode = ""
	r._ouCode = ""
	r._page = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.price.promotion.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterPromotionCode is OuterPromotionCode Setter
// 外部档期code
func (r *AlibabaPricePromotionActivityQueryAPIRequest) SetOuterPromotionCode(_outerPromotionCode string) error {
	r._outerPromotionCode = _outerPromotionCode
	r.Set("outer_promotion_code", _outerPromotionCode)
	return nil
}

// GetOuterPromotionCode OuterPromotionCode Getter
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetOuterPromotionCode() string {
	return r._outerPromotionCode
}

// SetOuCode is OuCode Setter
// TOB店仓编码
func (r *AlibabaPricePromotionActivityQueryAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetOuCode() string {
	return r._ouCode
}

// SetPage is Page Setter
// 页码
func (r *AlibabaPricePromotionActivityQueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 页码大小
func (r *AlibabaPricePromotionActivityQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaPricePromotionActivityQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPricePromotionActivityQueryRequest()
	},
}

// GetAlibabaPricePromotionActivityQueryRequest 从 sync.Pool 获取 AlibabaPricePromotionActivityQueryAPIRequest
func GetAlibabaPricePromotionActivityQueryAPIRequest() *AlibabaPricePromotionActivityQueryAPIRequest {
	return poolAlibabaPricePromotionActivityQueryAPIRequest.Get().(*AlibabaPricePromotionActivityQueryAPIRequest)
}

// ReleaseAlibabaPricePromotionActivityQueryAPIRequest 将 AlibabaPricePromotionActivityQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaPricePromotionActivityQueryAPIRequest(v *AlibabaPricePromotionActivityQueryAPIRequest) {
	v.Reset()
	poolAlibabaPricePromotionActivityQueryAPIRequest.Put(v)
}
