package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderQuery 订单详情查询
// taobao.train.agent.order.query
//
// 订单详情查询接口
func TaobaoTrainAgentOrderQuery(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderQueryAPIRequest, session string) (*train.TaobaoTrainAgentOrderQueryAPIResponse, error) {
	var resp train.TaobaoTrainAgentOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
