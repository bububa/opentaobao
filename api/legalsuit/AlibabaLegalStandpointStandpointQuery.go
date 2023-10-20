package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointStandpointQuery 查询具体口径
// alibaba.legal.standpoint.standpoint.query
//
// 查询具体口径
func AlibabaLegalStandpointStandpointQuery(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointStandpointQueryAPIRequest, session string) (*legalsuit.AlibabaLegalStandpointStandpointQueryAPIResponse, error) {
	var resp legalsuit.AlibabaLegalStandpointStandpointQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
