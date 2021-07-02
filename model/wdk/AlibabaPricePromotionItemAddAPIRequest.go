package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionItemAddAPIRequest) GetApiMethodName() string {
	return "alibaba.price.promotion.item.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionItemAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
