package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemPromotionRuleGetAPIRequest
获取商品已生效营销活动更新规则 API请求
taobao.item.promotion.rule.get

获取商品已生效的更新规则信息，主要包含库存禁止修改，商品一口价禁止修改，库存减少锁定等规则生效信息 */
type TaobaoItemPromotionRuleGetAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
}

// New
