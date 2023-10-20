package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentTostationConfirm 线下票确认送票至车站服务
// taobao.train.agent.tostation.confirm
//
// 送票至车站的订单，代理商确认配送到站
func TaobaoTrainAgentTostationConfirm(clt *core.SDKClient, req *train.TaobaoTrainAgentTostationConfirmAPIRequest, resp *train.TaobaoTrainAgentTostationConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
