package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitSealPush 法宝推送用印
// alibaba.legal.suit.seal.push
//
// 法宝推送用印
func AlibabaLegalSuitSealPush(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitSealPushAPIRequest, resp *legalsuit.AlibabaLegalSuitSealPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
