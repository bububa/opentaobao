package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangInventoryBatchQuery 批量查询库存
// alibaba.dchain.aoxiang.inventory.batch.query
//
// 批量查询库存
func AlibabaDchainAoxiangInventoryBatchQuery(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangInventoryBatchQueryAPIRequest, resp *ascp.AlibabaDchainAoxiangInventoryBatchQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
