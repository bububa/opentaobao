package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelDistributorProductSelect 供应链渠道中心品的选品接口（淘外分销商专用）
// alibaba.ascp.channel.distributor.product.select
//
// 此api为淘外分销的品的选品标准api，淘外分销商专用
func AlibabaAscpChannelDistributorProductSelect(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelDistributorProductSelectAPIRequest, resp *ascpchannel.AlibabaAscpChannelDistributorProductSelectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
