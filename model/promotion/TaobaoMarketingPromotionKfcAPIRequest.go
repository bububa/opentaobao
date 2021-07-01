package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMarketingPromotionKfcAPIRequest
定向优惠活动名称与描述违禁词检查 API请求
taobao.marketing.promotion.kfc

活动名称与描述违禁词检查 */
type TaobaoMarketingPromotionKfcAPIRequest struct {
	model.Params
	// 活动名称
	_promotionTitle string
	// 活动描述
	_promotionDesc string
}

// New
