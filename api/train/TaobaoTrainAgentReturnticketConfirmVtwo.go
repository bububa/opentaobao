package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentreturnticketconfirmvtwo 退票通知
// taobao.train.agent.returnticket.confirm.vtwo
//
// 火车票代理商接口——退票通知回调
func Taobaotrainagentreturnticketconfirmvtwo(clt *core.SDKClient, req *train.TaobaotrainagentreturnticketconfirmvtwoAPIRequest, session string) (*train.TaobaotrainagentreturnticketconfirmvtwoAPIResponse, error) {
	var resp train.TaobaotrainagentreturnticketconfirmvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
