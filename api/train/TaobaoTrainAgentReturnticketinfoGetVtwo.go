package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentReturnticketinfoGetVtwo 代理商获取退票详情回调
// taobao.train.agent.returnticketinfo.get.vtwo
//
// 代理商获取退票详情回调
func TaobaoTrainAgentReturnticketinfoGetVtwo(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest, resp *train.TaobaoTrainAgentReturnticketinfoGetVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
