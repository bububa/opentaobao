package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// TaobaoAlitripIeAgentShoppingPush 国际机票大卖家Shopping推送
// taobao.alitrip.ie.agent.shopping.push
//
// 用于国际机票大卖家主动推送Shopping结果更新缓存报价。
func TaobaoAlitripIeAgentShoppingPush(clt *core.SDKClient, req *flight.TaobaoAlitripIeAgentShoppingPushAPIRequest, resp *flight.TaobaoAlitripIeAgentShoppingPushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
