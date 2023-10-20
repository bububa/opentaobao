package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderLock 锁单
// taobao.train.agent.order.lock
//
// 锁单
func TaobaoTrainAgentOrderLock(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderLockAPIRequest, resp *train.TaobaoTrainAgentOrderLockAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
