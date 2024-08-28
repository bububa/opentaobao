package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentGetRefund 代理商获取订单退票信息
// taobao.train.agent.get.refund
//
// 代理商获取订单信息回调API
func TaobaoTrainAgentGetRefund(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentGetRefundAPIRequest, resp *train.TaobaoTrainAgentGetRefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
