package film

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/film"
)

/* TaobaoFilmLotteryDraw
淘票票抽奖发放权益API
taobao.film.lottery.draw

对外第三方合作渠道通过抽奖形式发码 */
func TaobaoFilmLotteryDraw(clt *core.SDKClient, req *film.TaobaoFilmLotteryDrawAPIRequest, session string) (*film.TaobaoFilmLotteryDrawAPIResponse, error) {
	var resp film.TaobaoFilmLotteryDrawAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
