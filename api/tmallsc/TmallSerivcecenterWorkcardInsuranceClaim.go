package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallSerivcecenterWorkcardInsuranceClaim 保险理赔回传工单记录
// tmall.serivcecenter.workcard.insurance.claim
//
// 保险理赔回传工单记录
func TmallSerivcecenterWorkcardInsuranceClaim(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallSerivcecenterWorkcardInsuranceClaimAPIRequest, resp *tmallsc.TmallSerivcecenterWorkcardInsuranceClaimAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
