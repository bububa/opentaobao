package film

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/film"
)

// Taobaofilmlotterysendcode 淘票票外部直发券
// taobao.film.lottery.sendcode
//
// 淘票票外部直发券
func Taobaofilmlotterysendcode(clt *core.SDKClient, req *film.TaobaofilmlotterysendcodeAPIRequest, session string) (*film.TaobaofilmlotterysendcodeAPIResponse, error) {
	var resp film.TaobaofilmlotterysendcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
