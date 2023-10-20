package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentorderquery 订单详情查询
// taobao.train.agent.order.query
//
// 订单详情查询接口
func Taobaotrainagentorderquery(clt *core.SDKClient, req *train.TaobaotrainagentorderqueryAPIRequest, session string) (*train.TaobaotrainagentorderqueryAPIResponse, error) {
	var resp train.TaobaotrainagentorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
