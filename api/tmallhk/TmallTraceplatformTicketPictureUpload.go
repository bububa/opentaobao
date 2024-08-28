package tmallhk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallTraceplatformTicketPictureUpload 上传小票图片
// tmall.traceplatform.ticket.picture.upload
//
// uploadPicture
func TmallTraceplatformTicketPictureUpload(ctx context.Context, clt *core.SDKClient, req *tmallhk.TmallTraceplatformTicketPictureUploadAPIRequest, resp *tmallhk.TmallTraceplatformTicketPictureUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
