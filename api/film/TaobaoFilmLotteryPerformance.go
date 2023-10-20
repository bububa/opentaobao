package film

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/film"
)

// Taobaofilmlotteryperformance 淘票票履约发放权益
// taobao.film.lottery.performance
//
// 对外第三方合作渠道通过抽奖形式发放权益
func Taobaofilmlotteryperformance(clt *core.SDKClient, req *film.TaobaofilmlotteryperformanceAPIRequest, session string) (*film.TaobaofilmlotteryperformanceAPIResponse, error) {
	var resp film.TaobaofilmlotteryperformanceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
