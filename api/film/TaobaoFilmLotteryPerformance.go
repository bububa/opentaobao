package film

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/film"
)

// TaobaoFilmLotteryPerformance 淘票票履约发放权益
// taobao.film.lottery.performance
//
// 对外第三方合作渠道通过抽奖形式发放权益
func TaobaoFilmLotteryPerformance(ctx context.Context, clt *core.SDKClient, req *film.TaobaoFilmLotteryPerformanceAPIRequest, resp *film.TaobaoFilmLotteryPerformanceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
