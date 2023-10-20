package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingopenposdiscountcodecreateAPIRequest pos一物一码创建 API请求
// alibaba.wdk.marketing.open.pos.discount.code.create
//
// pos一物一码创建
type AlibabawdkmarketingopenposdiscountcodecreateAPIRequest struct {
	model.Params
	// 请求信息
	_uniqueDiscountCodeRequest *UniqueDiscountCodeRequest
}

// NewAlibabawdkmarketingopenposdiscountcodecreateRequest 初始化AlibabawdkmarketingopenposdiscountcodecreateAPIRequest对象
func NewAlibabawdkmarketingopenposdiscountcodecreateRequest() *AlibabawdkmarketingopenposdiscountcodecreateAPIRequest {
	return &AlibabawdkmarketingopenposdiscountcodecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingopenposdiscountcodecreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.pos.discount.code.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingopenposdiscountcodecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingopenposdiscountcodecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUniqueDiscountCodeRequest is UniqueDiscountCodeRequest Setter
// 请求信息
func (r *AlibabawdkmarketingopenposdiscountcodecreateAPIRequest) SetUniqueDiscountCodeRequest(_uniqueDiscountCodeRequest *UniqueDiscountCodeRequest) error {
	r._uniqueDiscountCodeRequest = _uniqueDiscountCodeRequest
	r.Set("unique_discount_code_request", _uniqueDiscountCodeRequest)
	return nil
}

// GetUniqueDiscountCodeRequest UniqueDiscountCodeRequest Getter
func (r AlibabawdkmarketingopenposdiscountcodecreateAPIRequest) GetUniqueDiscountCodeRequest() *UniqueDiscountCodeRequest {
	return r._uniqueDiscountCodeRequest
}
