package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaoopentradeactivityquery 查询尖货活动信息
// taobao.opentrade.activity.query
//
// 尖货交易活动信息配置，查询尖货活动信息
func Taobaoopentradeactivityquery(clt *core.SDKClient, req *opentrade.TaobaoopentradeactivityqueryAPIRequest, session string) (*opentrade.TaobaoopentradeactivityqueryAPIResponse, error) {
	var resp opentrade.TaobaoopentradeactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
