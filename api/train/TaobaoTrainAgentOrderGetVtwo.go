package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderGetVtwo 代理商获取订单信息回调APIv2--增加鉴权校验
// taobao.train.agent.order.get.vtwo
//
// 代理商获取订单信息回调API
func TaobaoTrainAgentOrderGetVtwo(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentOrderGetVtwoAPIRequest, resp *train.TaobaoTrainAgentOrderGetVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
