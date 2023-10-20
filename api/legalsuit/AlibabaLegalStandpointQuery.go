package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointQuery 口径查询
// alibaba.legal.standpoint.query
//
// 口径查询
func AlibabaLegalStandpointQuery(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointQueryAPIRequest, resp *legalsuit.AlibabaLegalStandpointQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
