package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketBeforeconsume 电子凭证验码前置确认
// taobao.vmarket.eticket.beforeconsume
//
// 商家验码之前的调用接口，用来同步到最新的订单状态并判断是否可以进行验码操作
func TaobaoVmarketEticketBeforeconsume(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoVmarketEticketBeforeconsumeAPIRequest, resp *eticket.TaobaoVmarketEticketBeforeconsumeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
