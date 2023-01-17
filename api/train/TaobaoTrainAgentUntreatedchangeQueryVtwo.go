package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentUntreatedchangeQueryVtwo 火车票代理商接口-查询待处理改签单列表-含鉴权校验
// taobao.train.agent.untreatedchange.query.vtwo
//
// 火车票代理商接口-查询待处理改签单列表-含鉴权校验
func TaobaoTrainAgentUntreatedchangeQueryVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentUntreatedchangeQueryVtwoAPIRequest, session string) (*train.TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse, error) {
	var resp train.TaobaoTrainAgentUntreatedchangeQueryVtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
