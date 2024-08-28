package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderLock 锁单
// taobao.train.agent.order.lock
//
// 锁单
func TaobaoTrainAgentOrderLock(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentOrderLockAPIRequest, resp *train.TaobaoTrainAgentOrderLockAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
