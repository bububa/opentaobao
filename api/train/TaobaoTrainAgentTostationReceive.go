package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagenttostationreceive 线下票送票至车站代理商确认用户已取票服务
// taobao.train.agent.tostation.receive
//
// 送票至车站的订单，代理商确认用户已取票
func Taobaotrainagenttostationreceive(clt *core.SDKClient, req *train.TaobaotrainagenttostationreceiveAPIRequest, session string) (*train.TaobaotrainagenttostationreceiveAPIResponse, error) {
	var resp train.TaobaotrainagenttostationreceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
