package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentreturnordersgetvtwo 获取待退票的订单v2--增加鉴权校验
// taobao.train.agent.returnorders.get.vtwo
//
// 代理商用来获取待退票的订单列表及数量，防止代理商掉单。
func Taobaotrainagentreturnordersgetvtwo(clt *core.SDKClient, req *train.TaobaotrainagentreturnordersgetvtwoAPIRequest, session string) (*train.TaobaotrainagentreturnordersgetvtwoAPIResponse, error) {
	var resp train.TaobaotrainagentreturnordersgetvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
