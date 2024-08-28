package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalStanpointAccept 采纳口径
// alibaba.legal.stanpoint.accept
//
// 采纳口径
func AlibabaLegalStanpointAccept(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalStanpointAcceptAPIRequest, resp *legalsuit.AlibabaLegalStanpointAcceptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
