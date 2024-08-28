package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiIdGet 获取请求id
// alibaba.happytrip.taxi.id.get
//
// 获取订单号
func AlibabaHappytripTaxiIdGet(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiIdGetAPIRequest, resp *happytrip.AlibabaHappytripTaxiIdGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
