package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainStopAgentCallback 火车票车次停运信息商家回调
// taobao.train.stop.agent.callback
//
// 火车票车次停运信息商家回调接口
func TaobaoTrainStopAgentCallback(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainStopAgentCallbackAPIRequest, resp *train.TaobaoTrainStopAgentCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
