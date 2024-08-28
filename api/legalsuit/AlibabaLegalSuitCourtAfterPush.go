package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitCourtAfterPush 更新或者新增庭后信息
// alibaba.legal.suit.court.after.push
//
// 供外部ISV供应商 推送庭后信息给集团诉讼
func AlibabaLegalSuitCourtAfterPush(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitCourtAfterPushAPIRequest, resp *legalsuit.AlibabaLegalSuitCourtAfterPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
