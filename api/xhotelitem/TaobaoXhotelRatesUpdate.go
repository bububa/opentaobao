package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRatesUpdate 价格推送接口（批量全量）
// taobao.xhotel.rates.update
//
// 酒店产品库rate批量更新房态信息
func TaobaoXhotelRatesUpdate(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRatesUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelRatesUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
