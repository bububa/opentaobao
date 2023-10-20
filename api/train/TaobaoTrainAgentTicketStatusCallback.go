package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentticketstatuscallback 代理商票状态查询回调
// taobao.train.agent.ticket.status.callback
//
// 代理商票状态查询结果回调
func Taobaotrainagentticketstatuscallback(clt *core.SDKClient, req *train.TaobaotrainagentticketstatuscallbackAPIRequest, session string) (*train.TaobaotrainagentticketstatuscallbackAPIResponse, error) {
	var resp train.TaobaotrainagentticketstatuscallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
