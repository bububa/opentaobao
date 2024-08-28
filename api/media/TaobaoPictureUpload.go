package media

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPictureUpload 上传单张图片
// taobao.picture.upload
//
// 图片空间上传接口
func TaobaoPictureUpload(ctx context.Context, clt *core.SDKClient, req *media.TaobaoPictureUploadAPIRequest, resp *media.TaobaoPictureUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
