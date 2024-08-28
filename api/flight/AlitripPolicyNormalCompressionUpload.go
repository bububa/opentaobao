package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyNormalCompressionUpload 大批量上传普通类型的单程/往返政策
// alitrip.policy.normal.compression.upload
//
// 大批量上传普通类型的单程/往返政策
func AlitripPolicyNormalCompressionUpload(ctx context.Context, clt *core.SDKClient, req *flight.AlitripPolicyNormalCompressionUploadAPIRequest, resp *flight.AlitripPolicyNormalCompressionUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
