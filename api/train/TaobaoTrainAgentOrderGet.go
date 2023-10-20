package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentorderget 代理商获取订单信息回调API
// taobao.train.agent.order.get
//
// 代理商获取订单信息回调API
func Taobaotrainagentorderget(clt *core.SDKClient, req *train.TaobaotrainagentordergetAPIRequest, session string) (*train.TaobaotrainagentordergetAPIResponse, error) {
	var resp train.TaobaotrainagentordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
