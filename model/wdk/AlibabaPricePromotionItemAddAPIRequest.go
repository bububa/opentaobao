package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPricePromotionItemAddAPIRequest
新增档期商品 API请求
alibaba.price.promotion.item.add

批量新增档期活动商品 */
type AlibabaPricePromotionItemAddAPIRequest struct {
	model.Params
	// 入参
	_promotionContent *PromotionContent
}

// New
