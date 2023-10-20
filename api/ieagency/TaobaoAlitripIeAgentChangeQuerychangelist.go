package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentChangeQuerychangelist 卖家查询改签列表
// taobao.alitrip.ie.agent.change.querychangelist
//
// 提供B2B卖家查看改签列表服务
func TaobaoAlitripIeAgentChangeQuerychangelist(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest, resp *ieagency.TaobaoAlitripIeAgentChangeQuerychangelistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
