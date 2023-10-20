package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Alibabatjbpictureupload 淘特图片空间上传单张图片
// alibaba.tjb.picture.upload
//
// 淘特图片空间上传单张图片
func Alibabatjbpictureupload(clt *core.SDKClient, req *media.AlibabatjbpictureuploadAPIRequest, session string) (*media.AlibabatjbpictureuploadAPIResponse, error) {
	var resp media.AlibabatjbpictureuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
