package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentExpressSet 线下票回填物流信息
// taobao.train.agent.express.set
//
// 线下票回填物流信息服务
func TaobaoTrainAgentExpressSet(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentExpressSetAPIRequest, resp *train.TaobaoTrainAgentExpressSetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
