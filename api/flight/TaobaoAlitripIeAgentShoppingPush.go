package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

/* TaobaoAlitripIeAgentShoppingPush
国际机票大卖家Shopping推送
taobao.alitrip.ie.agent.shopping.push

用于国际机票大卖家主动推送Shopping结果更新缓存报价。 */
func TaobaoAlitripIeAgentShoppingPush(clt *core.SDKClient, req *flight.TaobaoAlitripIeAgentShoppingPushAPIRequest, session string) (*flight.TaobaoAlitripIeAgentShoppingPushAPIResponse, error) {
	var resp flight.TaobaoAlitripIeAgentShoppingPushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
