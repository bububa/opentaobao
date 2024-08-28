package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentExpressSetVtwo 线下票回填物流信息v2--增加鉴权校验
// taobao.train.agent.express.set.vtwo
//
// 线下票回填物流信息服务
func TaobaoTrainAgentExpressSetVtwo(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentExpressSetVtwoAPIRequest, resp *train.TaobaoTrainAgentExpressSetVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
