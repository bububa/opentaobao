package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointQuery 口径查询
// alibaba.legal.standpoint.query
//
// 口径查询
func AlibabaLegalStandpointQuery(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointQueryAPIRequest, resp *legalsuit.AlibabaLegalStandpointQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
