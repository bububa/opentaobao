package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaTjbPictureFolderCreate 淘特图片空间创建文件夹
// alibaba.tjb.picture.folder.create
//
// 淘特图片空间创建文件夹
func AlibabaTjbPictureFolderCreate(clt *core.SDKClient, req *media.AlibabaTjbPictureFolderCreateAPIRequest, resp *media.AlibabaTjbPictureFolderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
