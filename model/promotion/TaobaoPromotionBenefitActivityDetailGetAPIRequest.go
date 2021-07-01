package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionBenefitActivityDetailGetAPIRequest
活动关联的权益详情获取 API请求
taobao.promotion.benefit.activity.detail.get

活动关联的权益详情获取 */
type TaobaoPromotionBenefitActivityDetailGetAPIRequest struct {
	model.Params
	// 查询活动关联权益详情的请求
	_queryRequest *ActivityRelationDetailRequest
}

// New
