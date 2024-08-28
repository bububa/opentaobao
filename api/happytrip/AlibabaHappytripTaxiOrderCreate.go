package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderCreate 用户叫车
// alibaba.happytrip.taxi.order.create
//
// 用户根据需要发起叫车请求，在发起请求之前必须事先获得order id.
func AlibabaHappytripTaxiOrderCreate(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderCreateAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
