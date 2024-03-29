package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelDistributorProductDetail 获取供应链渠道中心品的详情接口（淘外分销商专用）
// alibaba.ascp.channel.distributor.product.detail
//
// 此api为淘外分销的品批量查询标准api，淘外分销商专用
func AlibabaAscpChannelDistributorProductDetail(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelDistributorProductDetailAPIRequest, resp *ascpchannel.AlibabaAscpChannelDistributorProductDetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
