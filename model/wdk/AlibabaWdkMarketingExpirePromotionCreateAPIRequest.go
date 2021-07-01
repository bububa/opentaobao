package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingExpirePromotionCreateAPIRequest
短保优惠创建 API请求
alibaba.wdk.marketing.expire.promotion.create

过期优惠优惠信息录入 */
type AlibabaWdkMarketingExpirePromotionCreateAPIRequest struct {
	model.Params
	// 创建短保优惠
	_param0 *ExpirePromotionBo
}

// New
