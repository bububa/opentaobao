package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentChangeHoldseatConfirm 火车票代理商接口——确认改签占座是否成功
// taobao.train.agent.change.holdseat.confirm
//
// 火车票代理商接口——确认改签占座是否成功
func TaobaoTrainAgentChangeHoldseatConfirm(clt *core.SDKClient, req *train.TaobaoTrainAgentChangeHoldseatConfirmAPIRequest, session string) (*train.TaobaoTrainAgentChangeHoldseatConfirmAPIResponse, error) {
	var resp train.TaobaoTrainAgentChangeHoldseatConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
