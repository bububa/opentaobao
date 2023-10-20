package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicySpecialCompressionUpload 大批量上传特殊类型的单程/往返政策
// alitrip.policy.special.compression.upload
//
// 大批量上传特殊类型的单程/往返政策
func AlitripPolicySpecialCompressionUpload(clt *core.SDKClient, req *flight.AlitripPolicySpecialCompressionUploadAPIRequest, session string) (*flight.AlitripPolicySpecialCompressionUploadAPIResponse, error) {
	var resp flight.AlitripPolicySpecialCompressionUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
