package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusSeatpriceGet 汽车票余票接口
// taobao.bus.seatprice.get
//
// 提供给商家，查询汽车票班次余票
func TaobaoBusSeatpriceGet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusSeatpriceGetAPIRequest, resp *bus.TaobaoBusSeatpriceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
