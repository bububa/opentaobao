package nlife

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cTradestatusDrive b2c订单状态驱动
// alibaba.nlife.b2c.tradestatus.drive
//
// 用于驱动零售+订单状态
func AlibabaNlifeB2cTradestatusDrive(ctx context.Context, clt *core.SDKClient, req *nlife.AlibabaNlifeB2cTradestatusDriveAPIRequest, resp *nlife.AlibabaNlifeB2cTradestatusDriveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
