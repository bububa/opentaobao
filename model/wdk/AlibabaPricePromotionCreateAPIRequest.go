package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPricePromotionCreateAPIRequest
营销档期活动创建 API请求
alibaba.price.promotion.create

大润发-盒马帮提供新增创建营销活动 */
type AlibabaPricePromotionCreateAPIRequest struct {
	model.Params
	// 档期活动参数
	_promotionActivityDo *PromotionActivityDo
}

// New
