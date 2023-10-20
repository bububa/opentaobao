package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangConsignorderBatchQuery 发货单批量查询
// alibaba.dchain.aoxiang.consignorder.batch.query
//
// 发货单批量查询
func AlibabaDchainAoxiangConsignorderBatchQuery(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest, resp *ascp.AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
