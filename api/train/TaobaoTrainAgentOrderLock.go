package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderLock 锁单
// taobao.train.agent.order.lock
//
// 锁单
func TaobaoTrainAgentOrderLock(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderLockAPIRequest, session string) (*train.TaobaoTrainAgentOrderLockAPIResponse, error) {
	var resp train.TaobaoTrainAgentOrderLockAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
