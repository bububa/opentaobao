package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderGet 订单详情
// alibaba.happytrip.taxi.order.get
//
// 获取订单状态及详情
func AlibabaHappytripTaxiOrderGet(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderGetAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
