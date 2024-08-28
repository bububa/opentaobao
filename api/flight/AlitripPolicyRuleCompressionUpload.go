package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyRuleCompressionUpload 大批量上传规则类型的单程/往返政策
// alitrip.policy.rule.compression.upload
//
// 大批量上传规则类型的单程/往返政策
func AlitripPolicyRuleCompressionUpload(ctx context.Context, clt *core.SDKClient, req *flight.AlitripPolicyRuleCompressionUploadAPIRequest, resp *flight.AlitripPolicyRuleCompressionUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
