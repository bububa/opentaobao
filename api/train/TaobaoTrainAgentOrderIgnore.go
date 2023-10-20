package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderIgnore 忽略订单
// taobao.train.agent.order.ignore
//
// 忽略订单
func TaobaoTrainAgentOrderIgnore(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderIgnoreAPIRequest, session string) (*train.TaobaoTrainAgentOrderIgnoreAPIResponse, error) {
	var resp train.TaobaoTrainAgentOrderIgnoreAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
