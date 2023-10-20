package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionBenefitActivityRelation 关联活动权益
// taobao.promotion.benefit.activity.relation
//
// 卖家活动中需要通过该API来关联的对应的权益。
func TaobaoPromotionBenefitActivityRelation(clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivityRelationAPIRequest, resp *promotion.TaobaoPromotionBenefitActivityRelationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
