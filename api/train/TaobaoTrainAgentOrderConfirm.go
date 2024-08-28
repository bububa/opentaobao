package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderConfirm 确认出票
// taobao.train.agent.order.confirm
//
// 确认出票
func TaobaoTrainAgentOrderConfirm(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentOrderConfirmAPIRequest, resp *train.TaobaoTrainAgentOrderConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
