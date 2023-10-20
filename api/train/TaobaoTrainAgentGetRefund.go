package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentgetrefund 代理商获取订单退票信息
// taobao.train.agent.get.refund
//
// 代理商获取订单信息回调API
func Taobaotrainagentgetrefund(clt *core.SDKClient, req *train.TaobaotrainagentgetrefundAPIRequest, session string) (*train.TaobaotrainagentgetrefundAPIResponse, error) {
	var resp train.TaobaotrainagentgetrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
