package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitCourtBeforePush 更新或保存庭前信息
// alibaba.legal.suit.court.before.push
//
// 更新或者保存庭前信息
func AlibabaLegalSuitCourtBeforePush(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitCourtBeforePushAPIRequest, resp *legalsuit.AlibabaLegalSuitCourtBeforePushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
