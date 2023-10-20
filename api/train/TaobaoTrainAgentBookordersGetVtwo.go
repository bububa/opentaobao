package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentbookordersgetvtwo 代理商获取待出票订单列表v2--增加鉴权校验
// taobao.train.agent.bookorders.get.vtwo
//
// 代理商获取待出票订单列表，只返回订单号
func Taobaotrainagentbookordersgetvtwo(clt *core.SDKClient, req *train.TaobaotrainagentbookordersgetvtwoAPIRequest, session string) (*train.TaobaotrainagentbookordersgetvtwoAPIResponse, error) {
	var resp train.TaobaotrainagentbookordersgetvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
