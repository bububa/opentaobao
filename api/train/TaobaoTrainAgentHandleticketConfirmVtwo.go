package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentHandleticketConfirmVtwo 代理商出票中v2--增加鉴权校验
// taobao.train.agent.handleticket.confirm.vtwo
//
// 代理商出票中
func TaobaoTrainAgentHandleticketConfirmVtwo(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest, resp *train.TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
