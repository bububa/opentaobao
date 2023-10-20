package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallSerivcecenterWorkcardInsuranceClaim 保险理赔回传工单记录
// tmall.serivcecenter.workcard.insurance.claim
//
// 保险理赔回传工单记录
func TmallSerivcecenterWorkcardInsuranceClaim(clt *core.SDKClient, req *tmallsc.TmallSerivcecenterWorkcardInsuranceClaimAPIRequest, resp *tmallsc.TmallSerivcecenterWorkcardInsuranceClaimAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
