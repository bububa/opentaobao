package retail

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// AlibabaRetailDeviceRoadStatusReset 贩卖机货道解锁
// alibaba.retail.device.road.status.reset
//
// 贩卖机货道解锁
func AlibabaRetailDeviceRoadStatusReset(ctx context.Context, clt *core.SDKClient, req *retail.AlibabaRetailDeviceRoadStatusResetAPIRequest, resp *retail.AlibabaRetailDeviceRoadStatusResetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
