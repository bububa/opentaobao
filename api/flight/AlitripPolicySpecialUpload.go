package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicySpecialUpload 特殊政策上传
// alitrip.policy.special.upload
//
// 上传特殊类型的单程/往返政策
func AlitripPolicySpecialUpload(clt *core.SDKClient, req *flight.AlitripPolicySpecialUploadAPIRequest, session string) (*flight.AlitripPolicySpecialUploadAPIResponse, error) {
	var resp flight.AlitripPolicySpecialUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
