package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketCodesGet 电子凭证码列表查询
// taobao.vmarket.eticket.codes.get
//
// 查询某个订单的所有码的列表
func TaobaoVmarketEticketCodesGet(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoVmarketEticketCodesGetAPIRequest, resp *eticket.TaobaoVmarketEticketCodesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
