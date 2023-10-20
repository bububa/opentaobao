package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentorderconfirm 确认出票
// taobao.train.agent.order.confirm
//
// 确认出票
func Taobaotrainagentorderconfirm(clt *core.SDKClient, req *train.TaobaotrainagentorderconfirmAPIRequest, session string) (*train.TaobaotrainagentorderconfirmAPIResponse, error) {
	var resp train.TaobaotrainagentorderconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
