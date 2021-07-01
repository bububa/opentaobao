package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

/* TaobaoAlitripIeAgentOrderHk
【国际机票】手工预定回填PNR
taobao.alitrip.ie.agent.order.hk

代理商通过手工预定PNR，并回填。 */
func TaobaoAlitripIeAgentOrderHk(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentOrderHkAPIRequest, session string) (*ieagency.TaobaoAlitripIeAgentOrderHkAPIResponse, error) {
	var resp ieagency.TaobaoAlitripIeAgentOrderHkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
