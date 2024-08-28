package tmallhk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallTraceplatformAwdcInfoUpload AWDC提交溯源信息
// tmall.traceplatform.awdc.info.upload
//
// 天猫溯源-AWDC-上传溯源信息
func TmallTraceplatformAwdcInfoUpload(ctx context.Context, clt *core.SDKClient, req *tmallhk.TmallTraceplatformAwdcInfoUploadAPIRequest, resp *tmallhk.TmallTraceplatformAwdcInfoUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
