package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// TaobaoTradeDrugOrderGet 查看订单详情
// taobao.trade.drug.order.get
//
// 商家查看订单详情
func TaobaoTradeDrugOrderGet(ctx context.Context, clt *core.SDKClient, req *alihealth2.TaobaoTradeDrugOrderGetAPIRequest, resp *alihealth2.TaobaoTradeDrugOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
