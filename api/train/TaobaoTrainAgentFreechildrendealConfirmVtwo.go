package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentFreechildrendealConfirmVtwo 免费儿童处理
// taobao.train.agent.freechildrendeal.confirm.vtwo
//
// 免费儿童列表查询
func TaobaoTrainAgentFreechildrendealConfirmVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentFreechildrendealConfirmVtwoAPIRequest, resp *train.TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
