package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

/* TaobaoPictureUpload
上传单张图片
taobao.picture.upload

图片空间上传接口 */
func TaobaoPictureUpload(clt *core.SDKClient, req *media.TaobaoPictureUploadAPIRequest, session string) (*media.TaobaoPictureUploadAPIResponse, error) {
	var resp media.TaobaoPictureUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
