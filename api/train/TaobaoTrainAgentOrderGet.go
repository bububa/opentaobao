package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderGet 代理商获取订单信息回调API
// taobao.train.agent.order.get
//
// 代理商获取订单信息回调API
func TaobaoTrainAgentOrderGet(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentOrderGetAPIRequest, resp *train.TaobaoTrainAgentOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
