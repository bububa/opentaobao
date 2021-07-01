package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPricePromotionActivityDeleteAPIRequest
删除档期活动 API请求
alibaba.price.promotion.activity.delete

删除盒马帮档期活动 */
type AlibabaPricePromotionActivityDeleteAPIRequest struct {
	model.Params
	// 外部主键
	_outerPromotionCode string
	// 经营店OU
	_ouCode string
}

// New
