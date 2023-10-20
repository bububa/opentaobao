package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentOrderConfirm 确认出票
// taobao.train.agent.order.confirm
//
// 确认出票
func TaobaoTrainAgentOrderConfirm(clt *core.SDKClient, req *train.TaobaoTrainAgentOrderConfirmAPIRequest, session string) (*train.TaobaoTrainAgentOrderConfirmAPIResponse, error) {
	var resp train.TaobaoTrainAgentOrderConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
