package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionBatchCancel 取消商品分销
// alibaba.dchain.aoxiang.item.distribution.batch.cancel
//
// 取消商品分销
func AlibabaDchainAoxiangItemDistributionBatchCancel(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest, resp *ascp.AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
