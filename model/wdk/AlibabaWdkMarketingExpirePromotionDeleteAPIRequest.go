package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingExpirePromotionDeleteAPIRequest
短保优惠删除 API请求
alibaba.wdk.marketing.expire.promotion.delete

短保优惠删除 */
type AlibabaWdkMarketingExpirePromotionDeleteAPIRequest struct {
	model.Params
	// 删除短保优惠
	_param0 *ExpirePromotionBo
}

// New
