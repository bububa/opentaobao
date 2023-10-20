package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyRuleCompressionUpload 大批量上传规则类型的单程/往返政策
// alitrip.policy.rule.compression.upload
//
// 大批量上传规则类型的单程/往返政策
func AlitripPolicyRuleCompressionUpload(clt *core.SDKClient, req *flight.AlitripPolicyRuleCompressionUploadAPIRequest, resp *flight.AlitripPolicyRuleCompressionUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
