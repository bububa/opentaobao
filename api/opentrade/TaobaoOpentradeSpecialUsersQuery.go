package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaoopentradespecialusersquery 专属下单标记信息查询
// taobao.opentrade.special.users.query
//
// 专属下单标记信息查询
func Taobaoopentradespecialusersquery(clt *core.SDKClient, req *opentrade.TaobaoopentradespecialusersqueryAPIRequest, session string) (*opentrade.TaobaoopentradespecialusersqueryAPIResponse, error) {
	var resp opentrade.TaobaoopentradespecialusersqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
