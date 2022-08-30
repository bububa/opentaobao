package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangConsignorderBatchQuery 发货单批量查询
// alibaba.dchain.aoxiang.consignorder.batch.query
//
// 发货单批量查询
func AlibabaDchainAoxiangConsignorderBatchQuery(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangConsignorderBatchQueryAPIRequest, session string) (*ascp.AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangConsignorderBatchQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
