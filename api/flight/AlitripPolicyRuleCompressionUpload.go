package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyRuleCompressionUpload 大批量上传规则类型的单程/往返政策
// alitrip.policy.rule.compression.upload
//
// 大批量上传规则类型的单程/往返政策
func AlitripPolicyRuleCompressionUpload(clt *core.SDKClient, req *flight.AlitripPolicyRuleCompressionUploadAPIRequest, session string) (*flight.AlitripPolicyRuleCompressionUploadAPIResponse, error) {
	var resp flight.AlitripPolicyRuleCompressionUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
