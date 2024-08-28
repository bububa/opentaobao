package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderDestinationModify 修改目的地
// alibaba.happytrip.taxi.order.destination.modify
//
// 通知ISV修改订单信息
func AlibabaHappytripTaxiOrderDestinationModify(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderDestinationModifyAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderDestinationModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
