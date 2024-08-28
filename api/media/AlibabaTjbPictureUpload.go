package media

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaTjbPictureUpload 淘特图片空间上传单张图片
// alibaba.tjb.picture.upload
//
// 淘特图片空间上传单张图片
func AlibabaTjbPictureUpload(ctx context.Context, clt *core.SDKClient, req *media.AlibabaTjbPictureUploadAPIRequest, resp *media.AlibabaTjbPictureUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
