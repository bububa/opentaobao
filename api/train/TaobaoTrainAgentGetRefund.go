package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

/* TaobaoTrainAgentGetRefund
代理商获取订单退票信息
taobao.train.agent.get.refund

代理商获取订单信息回调API */
func TaobaoTrainAgentGetRefund(clt *core.SDKClient, req *train.TaobaoTrainAgentGetRefundAPIRequest, session string) (*train.TaobaoTrainAgentGetRefundAPIResponse, error) {
	var resp train.TaobaoTrainAgentGetRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
