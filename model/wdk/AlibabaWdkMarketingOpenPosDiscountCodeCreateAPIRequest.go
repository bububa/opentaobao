package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest pos一物一码创建 API请求
// alibaba.wdk.marketing.open.pos.discount.code.create
//
// pos一物一码创建
type AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest struct {
	model.Params
	// 请求信息
	_uniqueDiscountCodeRequest *UniqueDiscountCodeRequest
}

// NewAlibabaWdkMarketingOpenPosDiscountCodeCreateRequest 初始化AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest对象
func NewAlibabaWdkMarketingOpenPosDiscountCodeCreateRequest() *AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest {
	return &AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest) Reset() {
	r._uniqueDiscountCodeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.open.pos.discount.code.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUniqueDiscountCodeRequest is UniqueDiscountCodeRequest Setter
// 请求信息
func (r *AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest) SetUniqueDiscountCodeRequest(_uniqueDiscountCodeRequest *UniqueDiscountCodeRequest) error {
	r._uniqueDiscountCodeRequest = _uniqueDiscountCodeRequest
	r.Set("unique_discount_code_request", _uniqueDiscountCodeRequest)
	return nil
}

// GetUniqueDiscountCodeRequest UniqueDiscountCodeRequest Getter
func (r AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest) GetUniqueDiscountCodeRequest() *UniqueDiscountCodeRequest {
	return r._uniqueDiscountCodeRequest
}

var poolAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingOpenPosDiscountCodeCreateRequest()
	},
}

// GetAlibabaWdkMarketingOpenPosDiscountCodeCreateRequest 从 sync.Pool 获取 AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest
func GetAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest() *AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest {
	return poolAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest.Get().(*AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest)
}

// ReleaseAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest 将 AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest(v *AlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingOpenPosDiscountCodeCreateAPIRequest.Put(v)
}
