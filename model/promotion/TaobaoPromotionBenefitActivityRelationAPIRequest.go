package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionBenefitActivityRelationAPIRequest
关联活动权益 API请求
taobao.promotion.benefit.activity.relation

卖家活动中需要通过该API来关联的对应的权益。 */
type TaobaoPromotionBenefitActivityRelationAPIRequest struct {
	model.Params
	// 活动关联权益请求参数
	_relationRequest *RelationActivityBenefitRequest
}

// New
