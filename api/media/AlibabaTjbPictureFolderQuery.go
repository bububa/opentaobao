package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaTjbPictureFolderQuery 淘特图片空间用户文件夹查询
// alibaba.tjb.picture.folder.query
//
// 淘特图片空间用户文件夹查询，返回用户所有的文件夹。
func AlibabaTjbPictureFolderQuery(clt *core.SDKClient, req *media.AlibabaTjbPictureFolderQueryAPIRequest, resp *media.AlibabaTjbPictureFolderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
