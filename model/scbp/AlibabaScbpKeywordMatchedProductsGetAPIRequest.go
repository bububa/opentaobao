package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpkeywordmatchedproductsgetAPIRequest 查询和词匹配的推广产品 API请求
// alibaba.scbp.keyword.matched.products.get
//
// 查询和词匹配的推广产品
type AlibabascbpkeywordmatchedproductsgetAPIRequest struct {
	model.Params
	// 已购买的关键词
	_adKeyword string
}

// NewAlibabascbpkeywordmatchedproductsgetRequest 初始化AlibabascbpkeywordmatchedproductsgetAPIRequest对象
func NewAlibabascbpkeywordmatchedproductsgetRequest() *AlibabascbpkeywordmatchedproductsgetAPIRequest {
	return &AlibabascbpkeywordmatchedproductsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpkeywordmatchedproductsgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.keyword.matched.products.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpkeywordmatchedproductsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpkeywordmatchedproductsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdKeyword is AdKeyword Setter
// 已购买的关键词
func (r *AlibabascbpkeywordmatchedproductsgetAPIRequest) SetAdKeyword(_adKeyword string) error {
	r._adKeyword = _adKeyword
	r.Set("ad_keyword", _adKeyword)
	return nil
}

// GetAdKeyword AdKeyword Getter
func (r AlibabascbpkeywordmatchedproductsgetAPIRequest) GetAdKeyword() string {
	return r._adKeyword
}
