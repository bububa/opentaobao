package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentuntreatedchangequeryvtwo 火车票代理商接口-查询待处理改签单列表-含鉴权校验
// taobao.train.agent.untreatedchange.query.vtwo
//
// 火车票代理商接口-查询待处理改签单列表-含鉴权校验
func Taobaotrainagentuntreatedchangequeryvtwo(clt *core.SDKClient, req *train.TaobaotrainagentuntreatedchangequeryvtwoAPIRequest, session string) (*train.TaobaotrainagentuntreatedchangequeryvtwoAPIResponse, error) {
	var resp train.TaobaotrainagentuntreatedchangequeryvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
