package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderGet 代理商获取订单信息回调API
// taobao.train.agent.order.get
//
// 代理商获取订单信息回调API
func TaobaoTrainAgentOrderGet(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderGetAPIRequest, session string) (*train.TaobaoTrainAgentOrderGetAPIResponse, error) {
	var resp train.TaobaoTrainAgentOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
