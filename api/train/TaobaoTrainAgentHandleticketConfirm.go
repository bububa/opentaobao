package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentHandleticketConfirm 代理商出票中
// taobao.train.agent.handleticket.confirm
//
// 代理商出票中
func TaobaoTrainAgentHandleticketConfirm(clt *core.SDKClient, req *train.TaobaoTrainAgentHandleticketConfirmAPIRequest, session string) (*train.TaobaoTrainAgentHandleticketConfirmAPIResponse, error) {
	var resp train.TaobaoTrainAgentHandleticketConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
