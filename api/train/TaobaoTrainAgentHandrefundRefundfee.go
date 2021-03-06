package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentHandrefundRefundfee 代理商手动退款接口
// taobao.train.agent.handrefund.refundfee
//
// 火车票代理商手动退款接口
func TaobaoTrainAgentHandrefundRefundfee(clt *core.SDKClient, req *train.TaobaoTrainAgentHandrefundRefundfeeAPIRequest, session string) (*train.TaobaoTrainAgentHandrefundRefundfeeAPIResponse, error) {
	var resp train.TaobaoTrainAgentHandrefundRefundfeeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
