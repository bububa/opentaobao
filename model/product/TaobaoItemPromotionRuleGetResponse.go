package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取商品已生效营销活动更新规则 APIResponse
taobao.item.promotion.rule.get

获取商品已生效的更新规则信息，主要包含库存禁止修改，商品一口价禁止修改，库存减少锁定等规则生效信息
*/
type TaobaoItemPromotionRuleGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemPromotionRuleGetResponse `json:"taobao_item_promotion_rule_get_response,omitempty"`
}

type TaobaoItemPromotionRuleGetResponse struct {

    // 商品规则信息
    Rules   []ItemPromotionRule `json:"rules,omitempty"`

    // 商品是否命中更新规则
    Effec   bool `json:"effec,omitempty"`

}
