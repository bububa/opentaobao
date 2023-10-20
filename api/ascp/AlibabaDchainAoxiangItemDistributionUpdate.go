package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionUpdate 更新商品分销内容
// alibaba.dchain.aoxiang.item.distribution.update
//
// 更新商品分销内容
func AlibabaDchainAoxiangItemDistributionUpdate(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionUpdateAPIRequest, resp *ascp.AlibabaDchainAoxiangItemDistributionUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
