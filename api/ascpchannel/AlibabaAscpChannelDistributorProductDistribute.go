package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelDistributorProductDistribute 分销商基于渠道产品铺货到商品
// alibaba.ascp.channel.distributor.product.distribute
//
// 分销商基于渠道产品铺货到商品
func AlibabaAscpChannelDistributorProductDistribute(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelDistributorProductDistributeAPIRequest, resp *ascpchannel.AlibabaAscpChannelDistributorProductDistributeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
