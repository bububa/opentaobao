package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitDominationPush 更新或者保存管辖信息
// alibaba.legal.suit.domination.push
//
// ISV推送管辖信息到诉讼平台
func AlibabaLegalSuitDominationPush(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitDominationPushAPIRequest, resp *legalsuit.AlibabaLegalSuitDominationPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
