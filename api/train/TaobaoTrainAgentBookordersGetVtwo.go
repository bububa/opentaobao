package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentBookordersGetVtwo 代理商获取待出票订单列表v2--增加鉴权校验
// taobao.train.agent.bookorders.get.vtwo
//
// 代理商获取待出票订单列表，只返回订单号
func TaobaoTrainAgentBookordersGetVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentBookordersGetVtwoAPIRequest, resp *train.TaobaoTrainAgentBookordersGetVtwoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
