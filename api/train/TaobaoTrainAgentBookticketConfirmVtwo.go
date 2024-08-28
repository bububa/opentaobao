package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentBookticketConfirmVtwo 火车票代理商接口——确认出票是否成功v2--增加鉴权校验
// taobao.train.agent.bookticket.confirm.vtwo
//
// 火车票代理商接口——确认出票是否成功
func TaobaoTrainAgentBookticketConfirmVtwo(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentBookticketConfirmVtwoAPIRequest, resp *train.TaobaoTrainAgentBookticketConfirmVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
