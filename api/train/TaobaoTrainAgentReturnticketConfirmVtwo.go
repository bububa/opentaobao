package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentReturnticketConfirmVtwo 退票通知
// taobao.train.agent.returnticket.confirm.vtwo
//
// 火车票代理商接口——退票通知回调
func TaobaoTrainAgentReturnticketConfirmVtwo(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentReturnticketConfirmVtwoAPIRequest, resp *train.TaobaoTrainAgentReturnticketConfirmVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
