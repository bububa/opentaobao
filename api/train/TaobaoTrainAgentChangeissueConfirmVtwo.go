package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentChangeissueConfirmVtwo 火车票代理商接口-跑腿改签出票回填-含鉴权校验
// taobao.train.agent.changeissue.confirm.vtwo
//
// 火车票代理商接口-跑腿改签出票回填-含鉴权校验
func TaobaoTrainAgentChangeissueConfirmVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest, session string) (*train.TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse, error) {
	var resp train.TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
