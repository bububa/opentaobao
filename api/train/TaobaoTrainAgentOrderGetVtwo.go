package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderGetVtwo 代理商获取订单信息回调APIv2--增加鉴权校验
// taobao.train.agent.order.get.vtwo
//
// 代理商获取订单信息回调API
func TaobaoTrainAgentOrderGetVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderGetVtwoAPIRequest, resp *train.TaobaoTrainAgentOrderGetVtwoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
