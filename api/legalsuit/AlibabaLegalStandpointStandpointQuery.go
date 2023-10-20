package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointStandpointQuery 查询具体口径
// alibaba.legal.standpoint.standpoint.query
//
// 查询具体口径
func AlibabaLegalStandpointStandpointQuery(clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointStandpointQueryAPIRequest, resp *legalsuit.AlibabaLegalStandpointStandpointQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
