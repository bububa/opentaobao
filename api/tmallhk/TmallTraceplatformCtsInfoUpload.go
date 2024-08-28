package tmallhk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallTraceplatformCtsInfoUpload CTS提交溯源信息
// tmall.traceplatform.cts.info.upload
//
// cts上传溯源信息
func TmallTraceplatformCtsInfoUpload(ctx context.Context, clt *core.SDKClient, req *tmallhk.TmallTraceplatformCtsInfoUploadAPIRequest, resp *tmallhk.TmallTraceplatformCtsInfoUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
