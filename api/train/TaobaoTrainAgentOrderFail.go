package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderFail 出票失败
// taobao.train.agent.order.fail
//
// 出票失败
func TaobaoTrainAgentOrderFail(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderFailAPIRequest, session string) (*train.TaobaoTrainAgentOrderFailAPIResponse, error) {
	var resp train.TaobaoTrainAgentOrderFailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
