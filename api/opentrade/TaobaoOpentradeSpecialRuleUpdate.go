package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaoopentradespecialruleupdate 专属下单更新限购规则
// taobao.opentrade.special.rule.update
//
// 对于专属下单的交易场景更新限购规则
func Taobaoopentradespecialruleupdate(clt *core.SDKClient, req *opentrade.TaobaoopentradespecialruleupdateAPIRequest, session string) (*opentrade.TaobaoopentradespecialruleupdateAPIResponse, error) {
	var resp opentrade.TaobaoopentradespecialruleupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
