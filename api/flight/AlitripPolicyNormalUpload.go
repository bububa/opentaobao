package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyNormalUpload 普通政策上传
// alitrip.policy.normal.upload
//
// 上传普通类型的单程/往返政策
func AlitripPolicyNormalUpload(clt *core.SDKClient, req *flight.AlitripPolicyNormalUploadAPIRequest, session string) (*flight.AlitripPolicyNormalUploadAPIResponse, error) {
	var resp flight.AlitripPolicyNormalUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
