package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitJudgementPush 推送裁判登记信息给集团法务系统
// alibaba.legal.suit.judgement.push
//
// isv推送裁判登记信息给集团法务系统
func AlibabaLegalSuitJudgementPush(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitJudgementPushAPIRequest, resp *legalsuit.AlibabaLegalSuitJudgementPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
