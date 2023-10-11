package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentFreechildrenlistQueryVtwo 免费儿童列表查询
// taobao.train.agent.freechildrenlist.query.vtwo
//
// 免费儿童列表查询
func TaobaoTrainAgentFreechildrenlistQueryVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentFreechildrenlistQueryVtwoAPIRequest, session string) (*train.TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse, error) {
	var resp train.TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
