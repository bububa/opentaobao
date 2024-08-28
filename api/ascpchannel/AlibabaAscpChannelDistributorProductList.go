package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelDistributorProductList 供应链渠道中心淘外分销品批量查询(分销商专用)
// alibaba.ascp.channel.distributor.product.list
//
// 此api为淘外分销的品批量查询标准api，淘外分销商专用
func AlibabaAscpChannelDistributorProductList(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelDistributorProductListAPIRequest, resp *ascpchannel.AlibabaAscpChannelDistributorProductListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
