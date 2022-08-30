package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyRuleUpload 规则政策上传
// alitrip.policy.rule.upload
//
// 上传特殊类型的单程/往返政策
func AlitripPolicyRuleUpload(clt *core.SDKClient, req *flight.AlitripPolicyRuleUploadAPIRequest, session string) (*flight.AlitripPolicyRuleUploadAPIResponse, error) {
	var resp flight.AlitripPolicyRuleUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
