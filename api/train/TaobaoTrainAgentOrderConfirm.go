package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderConfirm 确认出票
// taobao.train.agent.order.confirm
//
// 确认出票
func TaobaoTrainAgentOrderConfirm(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderConfirmAPIRequest, resp *train.TaobaoTrainAgentOrderConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
