package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeSpecialRuleUpdateAPIRequest
专属下单更新限购规则 API请求
taobao.opentrade.special.rule.update

对于专属下单的交易场景更新限购规则 */
type TaobaoOpentradeSpecialRuleUpdateAPIRequest struct {
	model.Params
	// 最大限购数量
	_limitNum int64
	// 商品id列表
	_itemIds []int64
}

// New
