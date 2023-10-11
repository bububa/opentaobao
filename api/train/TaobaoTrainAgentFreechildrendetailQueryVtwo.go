package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentFreechildrendetailQueryVtwo 免费儿童详情
// taobao.train.agent.freechildrendetail.query.vtwo
//
// 免费儿童列表详情
func TaobaoTrainAgentFreechildrendetailQueryVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentFreechildrendetailQueryVtwoAPIRequest, session string) (*train.TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse, error) {
	var resp train.TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
