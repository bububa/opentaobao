package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketStoreGet 获取电子凭证预约门店信息
// taobao.vmarket.eticket.store.get
//
// 用于给外部商家查询电子凭证预约门店信息
func TaobaoVmarketEticketStoreGet(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoVmarketEticketStoreGetAPIRequest, resp *eticket.TaobaoVmarketEticketStoreGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
