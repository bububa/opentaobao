package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentReturnticketinfoGetVtwo 代理商获取退票详情回调
// taobao.train.agent.returnticketinfo.get.vtwo
//
// 代理商获取退票详情回调
func TaobaoTrainAgentReturnticketinfoGetVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest, session string) (*train.TaobaoTrainAgentReturnticketinfoGetVtwoAPIResponse, error) {
	var resp train.TaobaoTrainAgentReturnticketinfoGetVtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
