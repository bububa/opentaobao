package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

/* TaobaoTrainAgentHandleticketConfirmVtwo
代理商出票中v2--增加鉴权校验
taobao.train.agent.handleticket.confirm.vtwo

代理商出票中 */
func TaobaoTrainAgentHandleticketConfirmVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest, session string) (*train.TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse, error) {
	var resp train.TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
