package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品已生效营销活动更新规则 APIResponse
taobao.item.promotion.rule.get

获取商品已生效的更新规则信息，主要包含库存禁止修改，商品一口价禁止修改，库存减少锁定等规则生效信息
*/
type TaobaoItemPromotionRuleGetAPIResponse struct {
    model.CommonResponse
    TaobaoItemPromotionRuleGetResponse
}

type TaobaoItemPromotionRuleGetResponse struct {
    XMLName xml.Name `xml:"item_promotion_rule_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品规则信息
    
    Rules   []ItemPromotionRule `json:"rules,omitempty" xml:"rules>item_promotion_rule,omitempty"`
    
    
    // 商品是否命中更新规则
    
    Effec   bool `json:"effec,omitempty" xml:"effec,omitempty"`

    
}
