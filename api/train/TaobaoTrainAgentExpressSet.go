package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentExpressSet 线下票回填物流信息
// taobao.train.agent.express.set
//
// 线下票回填物流信息服务
func TaobaoTrainAgentExpressSet(clt *core.SDKClient, req *train.TaobaoTrainAgentExpressSetAPIRequest, session string) (*train.TaobaoTrainAgentExpressSetAPIResponse, error) {
	var resp train.TaobaoTrainAgentExpressSetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
