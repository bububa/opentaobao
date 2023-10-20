package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangDeliverytemplateQuery 商家运费模板查询
// alibaba.dchain.aoxiang.deliverytemplate.query
//
// 商家运费模板查询
func AlibabaDchainAoxiangDeliverytemplateQuery(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangDeliverytemplateQueryAPIRequest, session string) (*ascp.AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangDeliverytemplateQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
