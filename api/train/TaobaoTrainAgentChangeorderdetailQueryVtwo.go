package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentChangeorderdetailQueryVtwo 火车票代理商接口-查询跑腿改签订单详情-含鉴权校验
// taobao.train.agent.changeorderdetail.query.vtwo
//
// 火车票代理商接口-查询跑腿改签订单详情-含鉴权校验
func TaobaoTrainAgentChangeorderdetailQueryVtwo(clt *core.SDKClient, req *train.TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest, session string) (*train.TaobaoTrainAgentChangeorderdetailQueryVtwoAPIResponse, error) {
	var resp train.TaobaoTrainAgentChangeorderdetailQueryVtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
