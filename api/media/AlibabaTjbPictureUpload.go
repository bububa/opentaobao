package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaTjbPictureUpload 淘特图片空间上传单张图片
// alibaba.tjb.picture.upload
//
// 淘特图片空间上传单张图片
func AlibabaTjbPictureUpload(clt *core.SDKClient, req *media.AlibabaTjbPictureUploadAPIRequest, resp *media.AlibabaTjbPictureUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
