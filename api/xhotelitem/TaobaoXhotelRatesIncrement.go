package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRatesIncrement 价格推送接口（批量增量）
// taobao.xhotel.rates.increment
//
// Rate库存&amp;价格增量更新接口，用户仅需要更新Rate中发生变化的库存日历&amp;价格日历即可
func TaobaoXhotelRatesIncrement(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRatesIncrementAPIRequest, resp *xhotelitem.TaobaoXhotelRatesIncrementAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
