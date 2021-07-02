package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentChangeAgreeVtwo 代理商同意改签v2--增加鉴权校验
// taobao.train.agent.change.agree.vtwo
//
// 代理商同意改签接口服务
func TaobaoTrainAgentChangeAgreeVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentChangeAgreeVtwoAPIRequest, session string) (*train.TaobaoTrainAgentChangeAgreeVtwoAPIResponse, error) {
	var resp train.TaobaoTrainAgentChangeAgreeVtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
