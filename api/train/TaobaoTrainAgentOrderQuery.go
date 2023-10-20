package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderQuery 订单详情查询
// taobao.train.agent.order.query
//
// 订单详情查询接口
func TaobaoTrainAgentOrderQuery(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderQueryAPIRequest, resp *train.TaobaoTrainAgentOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
