package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaolifereservationtradeconsumenotice 生服购后预约外部核销
// taobao.life.reservation.trade.consume.notice
//
// 生服团购商品，购后预约。外部ISV进行核销
func Taobaolifereservationtradeconsumenotice(clt *core.SDKClient, req *trade.TaobaolifereservationtradeconsumenoticeAPIRequest, session string) (*trade.TaobaolifereservationtradeconsumenoticeAPIResponse, error) {
	var resp trade.TaobaolifereservationtradeconsumenoticeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
