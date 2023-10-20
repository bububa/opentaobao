package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentfreechildrendetailqueryvtwo 免费儿童详情
// taobao.train.agent.freechildrendetail.query.vtwo
//
// 免费儿童列表详情
func Taobaotrainagentfreechildrendetailqueryvtwo(clt *core.SDKClient, req *train.TaobaotrainagentfreechildrendetailqueryvtwoAPIRequest, session string) (*train.TaobaotrainagentfreechildrendetailqueryvtwoAPIResponse, error) {
	var resp train.TaobaotrainagentfreechildrendetailqueryvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
