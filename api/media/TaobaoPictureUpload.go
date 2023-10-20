package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Taobaopictureupload 上传单张图片
// taobao.picture.upload
//
// 图片空间上传接口
func Taobaopictureupload(clt *core.SDKClient, req *media.TaobaopictureuploadAPIRequest, session string) (*media.TaobaopictureuploadAPIResponse, error) {
	var resp media.TaobaopictureuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
