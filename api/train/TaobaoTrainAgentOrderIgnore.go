package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentorderignore 忽略订单
// taobao.train.agent.order.ignore
//
// 忽略订单
func Taobaotrainagentorderignore(clt *core.SDKClient, req *train.TaobaotrainagentorderignoreAPIRequest, session string) (*train.TaobaotrainagentorderignoreAPIResponse, error) {
	var resp train.TaobaotrainagentorderignoreAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
