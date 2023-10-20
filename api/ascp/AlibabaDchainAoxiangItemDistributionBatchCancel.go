package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionBatchCancel 取消商品分销
// alibaba.dchain.aoxiang.item.distribution.batch.cancel
//
// 取消商品分销
func AlibabaDchainAoxiangItemDistributionBatchCancel(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest, session string) (*ascp.AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
