package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalNewdraftstandpointQuery 未采纳口径查询(新)
// alibaba.legal.newdraftstandpoint.query
//
// 未采纳口径查询(新)
func AlibabaLegalNewdraftstandpointQuery(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalNewdraftstandpointQueryAPIRequest, resp *legalsuit.AlibabaLegalNewdraftstandpointQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
