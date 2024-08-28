package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketConsume 电子票券消费通知
// taobao.vmarket.eticket.consume
//
// 外部合作商家电子票券消费回调接口
func TaobaoVmarketEticketConsume(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoVmarketEticketConsumeAPIRequest, resp *eticket.TaobaoVmarketEticketConsumeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
