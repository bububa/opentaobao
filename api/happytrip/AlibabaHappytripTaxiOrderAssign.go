package happytrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderAssign 订单指派
// alibaba.happytrip.taxi.order.assign
//
// 通知供应商订单指派成功
func AlibabaHappytripTaxiOrderAssign(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderAssignAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderAssignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
