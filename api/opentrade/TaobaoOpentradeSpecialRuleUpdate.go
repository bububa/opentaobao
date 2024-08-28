package opentrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeSpecialRuleUpdate 专属下单更新限购规则
// taobao.opentrade.special.rule.update
//
// 对于专属下单的交易场景更新限购规则
func TaobaoOpentradeSpecialRuleUpdate(ctx context.Context, clt *core.SDKClient, req *opentrade.TaobaoOpentradeSpecialRuleUpdateAPIRequest, resp *opentrade.TaobaoOpentradeSpecialRuleUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
