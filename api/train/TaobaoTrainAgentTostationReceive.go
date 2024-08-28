package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentTostationReceive 线下票送票至车站代理商确认用户已取票服务
// taobao.train.agent.tostation.receive
//
// 送票至车站的订单，代理商确认用户已取票
func TaobaoTrainAgentTostationReceive(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentTostationReceiveAPIRequest, resp *train.TaobaoTrainAgentTostationReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
