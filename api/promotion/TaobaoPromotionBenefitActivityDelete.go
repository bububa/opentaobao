package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionBenefitActivityDelete 删除关联的活动权益
// taobao.promotion.benefit.activity.delete
//
// 删除关联的活动权益
func TaobaoPromotionBenefitActivityDelete(clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivityDeleteAPIRequest, resp *promotion.TaobaoPromotionBenefitActivityDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
