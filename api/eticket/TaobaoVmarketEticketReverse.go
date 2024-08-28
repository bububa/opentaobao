package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketReverse 电子凭证冲正接口
// taobao.vmarket.eticket.reverse
//
// 电子凭证平台冲正接口
func TaobaoVmarketEticketReverse(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoVmarketEticketReverseAPIRequest, resp *eticket.TaobaoVmarketEticketReverseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
