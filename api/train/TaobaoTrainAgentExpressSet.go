package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentexpressset 线下票回填物流信息
// taobao.train.agent.express.set
//
// 线下票回填物流信息服务
func Taobaotrainagentexpressset(clt *core.SDKClient, req *train.TaobaotrainagentexpresssetAPIRequest, session string) (*train.TaobaotrainagentexpresssetAPIResponse, error) {
	var resp train.TaobaotrainagentexpresssetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
