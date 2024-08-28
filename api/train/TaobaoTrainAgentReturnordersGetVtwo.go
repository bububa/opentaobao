package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentReturnordersGetVtwo 获取待退票的订单v2--增加鉴权校验
// taobao.train.agent.returnorders.get.vtwo
//
// 代理商用来获取待退票的订单列表及数量，防止代理商掉单。
func TaobaoTrainAgentReturnordersGetVtwo(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentReturnordersGetVtwoAPIRequest, resp *train.TaobaoTrainAgentReturnordersGetVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
