package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyNormalCompressionUpload 大批量上传普通类型的单程/往返政策
// alitrip.policy.normal.compression.upload
//
// 大批量上传普通类型的单程/往返政策
func AlitripPolicyNormalCompressionUpload(clt *core.SDKClient, req *flight.AlitripPolicyNormalCompressionUploadAPIRequest, session string) (*flight.AlitripPolicyNormalCompressionUploadAPIResponse, error) {
	var resp flight.AlitripPolicyNormalCompressionUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
