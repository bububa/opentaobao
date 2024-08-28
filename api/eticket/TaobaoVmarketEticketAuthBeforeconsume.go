package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketAuthBeforeconsume 核销放行的查询接口
// taobao.vmarket.eticket.auth.beforeconsume
//
// 针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销查询接口
func TaobaoVmarketEticketAuthBeforeconsume(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoVmarketEticketAuthBeforeconsumeAPIRequest, resp *eticket.TaobaoVmarketEticketAuthBeforeconsumeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
