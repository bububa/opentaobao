package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentchangeorderdetailqueryvtwo 火车票代理商接口-查询跑腿改签订单详情-含鉴权校验
// taobao.train.agent.changeorderdetail.query.vtwo
//
// 火车票代理商接口-查询跑腿改签订单详情-含鉴权校验
func Taobaotrainagentchangeorderdetailqueryvtwo(clt *core.SDKClient, req *train.TaobaotrainagentchangeorderdetailqueryvtwoAPIRequest, session string) (*train.TaobaotrainagentchangeorderdetailqueryvtwoAPIResponse, error) {
	var resp train.TaobaotrainagentchangeorderdetailqueryvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
