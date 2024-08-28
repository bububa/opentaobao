package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusNumbersStockpriceUpdate 汽车票更新价格库存
// taobao.bus.numbers.stockprice.update
//
// 用于汽车票代理商更新价格库存
func TaobaoBusNumbersStockpriceUpdate(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusNumbersStockpriceUpdateAPIRequest, resp *bus.TaobaoBusNumbersStockpriceUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
