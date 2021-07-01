package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

/* TaobaoPromotionBenefitActivityDelete
删除关联的活动权益
taobao.promotion.benefit.activity.delete

删除关联的活动权益 */
func TaobaoPromotionBenefitActivityDelete(clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivityDeleteAPIRequest, session string) (*promotion.TaobaoPromotionBenefitActivityDeleteAPIResponse, error) {
	var resp promotion.TaobaoPromotionBenefitActivityDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
