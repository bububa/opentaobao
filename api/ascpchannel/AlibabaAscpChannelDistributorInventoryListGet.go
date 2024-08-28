package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelDistributorInventoryListGet 批量查询渠道库存
// alibaba.ascp.channel.distributor.inventory.list.get
//
// 淘外分销批量查询渠道产品的库存
func AlibabaAscpChannelDistributorInventoryListGet(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelDistributorInventoryListGetAPIRequest, resp *ascpchannel.AlibabaAscpChannelDistributorInventoryListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
