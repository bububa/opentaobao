package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelDistributorInventoryGet 链渠道中心淘外库存查询
// alibaba.ascp.channel.distributor.inventory.get
//
// 此api为淘外分销的渠道产品库存查询标准api，淘外分销商专用
func AlibabaAscpChannelDistributorInventoryGet(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelDistributorInventoryGetAPIRequest, resp *ascpchannel.AlibabaAscpChannelDistributorInventoryGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
