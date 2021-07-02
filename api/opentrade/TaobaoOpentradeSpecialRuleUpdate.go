package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeSpecialRuleUpdate 专属下单更新限购规则
// taobao.opentrade.special.rule.update
//
// 对于专属下单的交易场景更新限购规则
func TaobaoOpentradeSpecialRuleUpdate(clt *core.SDKClient, req *opentrade.TaobaoOpentradeSpecialRuleUpdateAPIRequest, session string) (*opentrade.TaobaoOpentradeSpecialRuleUpdateAPIResponse, error) {
	var resp opentrade.TaobaoOpentradeSpecialRuleUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
