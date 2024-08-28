package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicySpecialUpload 特殊政策上传
// alitrip.policy.special.upload
//
// 上传特殊类型的单程/往返政策
func AlitripPolicySpecialUpload(ctx context.Context, clt *core.SDKClient, req *flight.AlitripPolicySpecialUploadAPIRequest, resp *flight.AlitripPolicySpecialUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
