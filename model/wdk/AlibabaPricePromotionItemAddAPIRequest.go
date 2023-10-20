package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPricePromotionItemAddAPIRequest 新增档期商品 API请求
// alibaba.price.promotion.item.add
//
// 批量新增档期活动商品
type AlibabaPricePromotionItemAddAPIRequest struct {
	model.Params
	// 入参
	_promotionContent *PromotionContent
}

// NewAlibabaPricePromotionItemAddRequest 初始化AlibabaPricePromotionItemAddAPIRequest对象
func NewAlibabaPricePromotionItemAddRequest() *AlibabaPricePromotionItemAddAPIRequest {
	return &AlibabaPricePromotionItemAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPricePromotionItemAddAPIRequest) Reset() {
	r._promotionContent = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionItemAddAPIRequest) GetApiMethodName() string {
	return "alibaba.price.promotion.item.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionItemAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPricePromotionItemAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPromotionContent is PromotionContent Setter
// 入参
func (r *AlibabaPricePromotionItemAddAPIRequest) SetPromotionContent(_promotionContent *PromotionContent) error {
	r._promotionContent = _promotionContent
	r.Set("promotion_content", _promotionContent)
	return nil
}

// GetPromotionContent PromotionContent Getter
func (r AlibabaPricePromotionItemAddAPIRequest) GetPromotionContent() *PromotionContent {
	return r._promotionContent
}

var poolAlibabaPricePromotionItemAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPricePromotionItemAddRequest()
	},
}

// GetAlibabaPricePromotionItemAddRequest 从 sync.Pool 获取 AlibabaPricePromotionItemAddAPIRequest
func GetAlibabaPricePromotionItemAddAPIRequest() *AlibabaPricePromotionItemAddAPIRequest {
	return poolAlibabaPricePromotionItemAddAPIRequest.Get().(*AlibabaPricePromotionItemAddAPIRequest)
}

// ReleaseAlibabaPricePromotionItemAddAPIRequest 将 AlibabaPricePromotionItemAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaPricePromotionItemAddAPIRequest(v *AlibabaPricePromotionItemAddAPIRequest) {
	v.Reset()
	poolAlibabaPricePromotionItemAddAPIRequest.Put(v)
}
