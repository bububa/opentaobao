package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TaobaoItemPromotionRuleGet
获取商品已生效营销活动更新规则
taobao.item.promotion.rule.get

获取商品已生效的更新规则信息，主要包含库存禁止修改，商品一口价禁止修改，库存减少锁定等规则生效信息 */
func TaobaoItemPromotionRuleGet(clt *core.SDKClient, req *product.TaobaoItemPromotionRuleGetAPIRequest, session string) (*product.TaobaoItemPromotionRuleGetAPIResponse, error) {
	var resp product.TaobaoItemPromotionRuleGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
