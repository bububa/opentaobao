package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderIgnore 忽略订单
// taobao.train.agent.order.ignore
//
// 忽略订单
func TaobaoTrainAgentOrderIgnore(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderIgnoreAPIRequest, resp *train.TaobaoTrainAgentOrderIgnoreAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
