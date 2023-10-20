package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentordergetvtwo 代理商获取订单信息回调APIv2--增加鉴权校验
// taobao.train.agent.order.get.vtwo
//
// 代理商获取订单信息回调API
func Taobaotrainagentordergetvtwo(clt *core.SDKClient, req *train.TaobaotrainagentordergetvtwoAPIRequest, session string) (*train.TaobaotrainagentordergetvtwoAPIResponse, error) {
	var resp train.TaobaotrainagentordergetvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
