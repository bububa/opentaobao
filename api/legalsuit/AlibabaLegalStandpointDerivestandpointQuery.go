package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointDerivestandpointQuery 查询衍生口径
// alibaba.legal.standpoint.derivestandpoint.query
//
// 查询衍生口径
func AlibabaLegalStandpointDerivestandpointQuery(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointDerivestandpointQueryAPIRequest, session string) (*legalsuit.AlibabaLegalStandpointDerivestandpointQueryAPIResponse, error) {
	var resp legalsuit.AlibabaLegalStandpointDerivestandpointQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
