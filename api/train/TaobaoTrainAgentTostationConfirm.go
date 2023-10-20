package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagenttostationconfirm 线下票确认送票至车站服务
// taobao.train.agent.tostation.confirm
//
// 送票至车站的订单，代理商确认配送到站
func Taobaotrainagenttostationconfirm(clt *core.SDKClient, req *train.TaobaotrainagenttostationconfirmAPIRequest, session string) (*train.TaobaotrainagenttostationconfirmAPIResponse, error) {
	var resp train.TaobaotrainagenttostationconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
