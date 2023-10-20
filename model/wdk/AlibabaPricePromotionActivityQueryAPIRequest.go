package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapricepromotionactivityqueryAPIRequest 查询盒马帮档期活动详情 API请求
// alibaba.price.promotion.activity.query
//
// 查询盒马帮档期活动详情
type AlibabapricepromotionactivityqueryAPIRequest struct {
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

// NewAlibabapricepromotionactivityqueryRequest 初始化AlibabapricepromotionactivityqueryAPIRequest对象
func NewAlibabapricepromotionactivityqueryRequest() *AlibabapricepromotionactivityqueryAPIRequest {
	return &AlibabapricepromotionactivityqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapricepromotionactivityqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.price.promotion.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapricepromotionactivityqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapricepromotionactivityqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterPromotionCode is OuterPromotionCode Setter
// 外部档期code
func (r *AlibabapricepromotionactivityqueryAPIRequest) SetOuterPromotionCode(_outerPromotionCode string) error {
	r._outerPromotionCode = _outerPromotionCode
	r.Set("outer_promotion_code", _outerPromotionCode)
	return nil
}

// GetOuterPromotionCode OuterPromotionCode Getter
func (r AlibabapricepromotionactivityqueryAPIRequest) GetOuterPromotionCode() string {
	return r._outerPromotionCode
}

// SetOuCode is OuCode Setter
// TOB店仓编码
func (r *AlibabapricepromotionactivityqueryAPIRequest) SetOuCode(_ouCode string) error {
	r._ouCode = _ouCode
	r.Set("ou_code", _ouCode)
	return nil
}

// GetOuCode OuCode Getter
func (r AlibabapricepromotionactivityqueryAPIRequest) GetOuCode() string {
	return r._ouCode
}

// SetPage is Page Setter
// 页码
func (r *AlibabapricepromotionactivityqueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabapricepromotionactivityqueryAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 页码大小
func (r *AlibabapricepromotionactivityqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabapricepromotionactivityqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
