package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// Tmalltraceplatformticketpictureupload 上传小票图片
// tmall.traceplatform.ticket.picture.upload
//
// uploadPicture
func Tmalltraceplatformticketpictureupload(clt *core.SDKClient, req *tmallhk.TmalltraceplatformticketpictureuploadAPIRequest, session string) (*tmallhk.TmalltraceplatformticketpictureuploadAPIResponse, error) {
	var resp tmallhk.TmalltraceplatformticketpictureuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
