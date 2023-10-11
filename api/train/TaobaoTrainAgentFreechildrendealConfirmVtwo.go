package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentFreechildrendealConfirmVtwo 免费儿童处理
// taobao.train.agent.freechildrendeal.confirm.vtwo
//
// 免费儿童列表查询
func TaobaoTrainAgentFreechildrendealConfirmVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentFreechildrendealConfirmVtwoAPIRequest, session string) (*train.TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse, error) {
	var resp train.TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
