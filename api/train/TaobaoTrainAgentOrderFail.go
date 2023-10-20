package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentorderfail 出票失败
// taobao.train.agent.order.fail
//
// 出票失败
func Taobaotrainagentorderfail(clt *core.SDKClient, req *train.TaobaotrainagentorderfailAPIRequest, session string) (*train.TaobaotrainagentorderfailAPIResponse, error) {
	var resp train.TaobaotrainagentorderfailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
