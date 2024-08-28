package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateUpload 协同单附件凭证上传
// alitrip.agent.coordinate.upload
//
// 协同单附件凭证上传
func AlitripAgentCoordinateUpload(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentCoordinateUploadAPIRequest, resp *flight.AlitripAgentCoordinateUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
