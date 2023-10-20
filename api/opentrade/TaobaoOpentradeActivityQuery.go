package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeActivityQuery 查询尖货活动信息
// taobao.opentrade.activity.query
//
// 尖货交易活动信息配置，查询尖货活动信息
func TaobaoOpentradeActivityQuery(clt *core.SDKClient, req *opentrade.TaobaoOpentradeActivityQueryAPIRequest, resp *opentrade.TaobaoOpentradeActivityQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
