package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Alibabatjbpicturefoldercreate 淘特图片空间创建文件夹
// alibaba.tjb.picture.folder.create
//
// 淘特图片空间创建文件夹
func Alibabatjbpicturefoldercreate(clt *core.SDKClient, req *media.AlibabatjbpicturefoldercreateAPIRequest, session string) (*media.AlibabatjbpicturefoldercreateAPIResponse, error) {
	var resp media.AlibabatjbpicturefoldercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
