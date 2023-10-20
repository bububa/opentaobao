package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentReturnticketConfirmVtwo 退票通知
// taobao.train.agent.returnticket.confirm.vtwo
//
// 火车票代理商接口——退票通知回调
func TaobaoTrainAgentReturnticketConfirmVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest, resp *train.TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
