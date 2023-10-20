package film

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/film"
)

// TaobaoFilmLotteryPerformance 淘票票履约发放权益
// taobao.film.lottery.performance
//
// 对外第三方合作渠道通过抽奖形式发放权益
func TaobaoFilmLotteryPerformance(clt *core.SDKClient, req *film.TaobaoFilmLotteryPerformanceAPIRequest, session string) (*film.TaobaoFilmLotteryPerformanceAPIResponse, error) {
	var resp film.TaobaoFilmLotteryPerformanceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
