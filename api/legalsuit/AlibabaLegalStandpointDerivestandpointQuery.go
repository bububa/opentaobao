package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointDerivestandpointQuery 查询衍生口径
// alibaba.legal.standpoint.derivestandpoint.query
//
// 查询衍生口径
func AlibabaLegalStandpointDerivestandpointQuery(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointDerivestandpointQueryAPIRequest, resp *legalsuit.AlibabaLegalStandpointDerivestandpointQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
