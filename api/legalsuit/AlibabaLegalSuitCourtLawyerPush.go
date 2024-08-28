package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitCourtLawyerPush 推荐律师接口
// alibaba.legal.suit.court.lawyer.push
//
// 为诉讼系统推荐律师
func AlibabaLegalSuitCourtLawyerPush(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitCourtLawyerPushAPIRequest, resp *legalsuit.AlibabaLegalSuitCourtLawyerPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
