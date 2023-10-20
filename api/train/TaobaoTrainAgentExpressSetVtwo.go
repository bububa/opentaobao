package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentExpressSetVtwo 线下票回填物流信息v2--增加鉴权校验
// taobao.train.agent.express.set.vtwo
//
// 线下票回填物流信息服务
func TaobaoTrainAgentExpressSetVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentExpressSetVtwoAPIRequest, resp *train.TaobaoTrainAgentExpressSetVtwoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
