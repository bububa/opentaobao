package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentorderlock 锁单
// taobao.train.agent.order.lock
//
// 锁单
func Taobaotrainagentorderlock(clt *core.SDKClient, req *train.TaobaotrainagentorderlockAPIRequest, session string) (*train.TaobaotrainagentorderlockAPIResponse, error) {
	var resp train.TaobaotrainagentorderlockAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
