package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyRuleUpload 规则政策上传
// alitrip.policy.rule.upload
//
// 上传特殊类型的单程/往返政策
func AlitripPolicyRuleUpload(ctx context.Context, clt *core.SDKClient, req *flight.AlitripPolicyRuleUploadAPIRequest, resp *flight.AlitripPolicyRuleUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
