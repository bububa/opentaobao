package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentChangeorderdetailQueryVtwo 火车票代理商接口-查询跑腿改签订单详情-含鉴权校验
// taobao.train.agent.changeorderdetail.query.vtwo
//
// 火车票代理商接口-查询跑腿改签订单详情-含鉴权校验
func TaobaoTrainAgentChangeorderdetailQueryVtwo(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentChangeorderdetailQueryVtwoAPIRequest, resp *train.TaobaoTrainAgentChangeorderdetailQueryVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
