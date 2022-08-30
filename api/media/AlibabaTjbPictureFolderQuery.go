package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaTjbPictureFolderQuery 淘特图片空间用户文件夹查询
// alibaba.tjb.picture.folder.query
//
// 淘特图片空间用户文件夹查询，返回用户所有的文件夹。
func AlibabaTjbPictureFolderQuery(clt *core.SDKClient, req *media.AlibabaTjbPictureFolderQueryAPIRequest, session string) (*media.AlibabaTjbPictureFolderQueryAPIResponse, error) {
	var resp media.AlibabaTjbPictureFolderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
