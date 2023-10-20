package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagenthandleticketconfirmvtwo 代理商出票中v2--增加鉴权校验
// taobao.train.agent.handleticket.confirm.vtwo
//
// 代理商出票中
func Taobaotrainagenthandleticketconfirmvtwo(clt *core.SDKClient, req *train.TaobaotrainagenthandleticketconfirmvtwoAPIRequest, session string) (*train.TaobaotrainagenthandleticketconfirmvtwoAPIResponse, error) {
	var resp train.TaobaotrainagenthandleticketconfirmvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
