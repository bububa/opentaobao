package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointStandpointQuery 查询具体口径
// alibaba.legal.standpoint.standpoint.query
//
// 查询具体口径
func AlibabaLegalStandpointStandpointQuery(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointStandpointQueryAPIRequest, resp *legalsuit.AlibabaLegalStandpointStandpointQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
