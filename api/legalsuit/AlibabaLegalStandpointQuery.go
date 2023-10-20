package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointQuery 口径查询
// alibaba.legal.standpoint.query
//
// 口径查询
func AlibabaLegalStandpointQuery(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointQueryAPIRequest, session string) (*legalsuit.AlibabaLegalStandpointQueryAPIResponse, error) {
	var resp legalsuit.AlibabaLegalStandpointQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
