package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantTbmaGet 码商查询淘宝码接口
// taobao.eticket.merchant.tbma.get
//
// 码商查询淘宝码接口
func TaobaoEticketMerchantTbmaGet(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoEticketMerchantTbmaGetAPIRequest, resp *eticket.TaobaoEticketMerchantTbmaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
