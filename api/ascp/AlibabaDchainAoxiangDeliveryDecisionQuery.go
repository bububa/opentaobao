package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangDeliveryDecisionQuery 查询黑白名单快递
// alibaba.dchain.aoxiang.delivery.decision.query
//
// 查询黑白名单快递
func AlibabaDchainAoxiangDeliveryDecisionQuery(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangDeliveryDecisionQueryAPIRequest, session string) (*ascp.AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
