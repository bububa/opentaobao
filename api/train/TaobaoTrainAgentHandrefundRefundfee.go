package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagenthandrefundrefundfee 代理商手动退款接口
// taobao.train.agent.handrefund.refundfee
//
// 火车票代理商手动退款接口
func Taobaotrainagenthandrefundrefundfee(clt *core.SDKClient, req *train.TaobaotrainagenthandrefundrefundfeeAPIRequest, session string) (*train.TaobaotrainagenthandrefundrefundfeeAPIResponse, error) {
	var resp train.TaobaotrainagenthandrefundrefundfeeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
