package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoPictureQnaigcUpload qnaigc业务线图片上传
// taobao.picture.qnaigc.upload
//
// qnaigc业务线图片上传
func TaobaoPictureQnaigcUpload(ctx context.Context, clt *core.SDKClient, req *util.TaobaoPictureQnaigcUploadAPIRequest, resp *util.TaobaoPictureQnaigcUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
