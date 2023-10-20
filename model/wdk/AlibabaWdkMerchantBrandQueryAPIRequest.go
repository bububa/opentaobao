package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmerchantbrandqueryAPIRequest 品牌查询接口 API请求
// alibaba.wdk.merchant.brand.query
//
// 三江erp对接时，提供品牌查询的接口
type AlibabawdkmerchantbrandqueryAPIRequest struct {
	model.Params
	// 关键词，不填就查全部
	_keyword string
	// 可不填，默认0
	_offset int64
	// 可不填，默认2000
	_pageSize int64
}

// NewAlibabawdkmerchantbrandqueryRequest 初始化AlibabawdkmerchantbrandqueryAPIRequest对象
func NewAlibabawdkmerchantbrandqueryRequest() *AlibabawdkmerchantbrandqueryAPIRequest {
	return &AlibabawdkmerchantbrandqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmerchantbrandqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.merchant.brand.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmerchantbrandqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmerchantbrandqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 关键词，不填就查全部
func (r *AlibabawdkmerchantbrandqueryAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabawdkmerchantbrandqueryAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetOffset is Offset Setter
// 可不填，默认0
func (r *AlibabawdkmerchantbrandqueryAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r AlibabawdkmerchantbrandqueryAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 可不填，默认2000
func (r *AlibabawdkmerchantbrandqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabawdkmerchantbrandqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
