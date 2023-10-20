package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPictureUpload 上传单张图片
// taobao.picture.upload
//
// 图片空间上传接口
func TaobaoPictureUpload(clt *core.SDKClient, req *media.TaobaoPictureUploadAPIRequest, resp *media.TaobaoPictureUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
