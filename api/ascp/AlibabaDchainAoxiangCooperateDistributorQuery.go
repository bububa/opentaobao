package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangCooperateDistributorQuery 商家关系查询分销商
// alibaba.dchain.aoxiang.cooperate.distributor.query
//
// 商家关系查询分销商
func AlibabaDchainAoxiangCooperateDistributorQuery(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangCooperateDistributorQueryAPIRequest, session string) (*ascp.AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangCooperateDistributorQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
