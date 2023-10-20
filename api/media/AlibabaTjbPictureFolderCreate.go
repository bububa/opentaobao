package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaTjbPictureFolderCreate 淘特图片空间创建文件夹
// alibaba.tjb.picture.folder.create
//
// 淘特图片空间创建文件夹
func AlibabaTjbPictureFolderCreate(clt *core.SDKClient, req *media.AlibabaTjbPictureFolderCreateAPIRequest, session string) (*media.AlibabaTjbPictureFolderCreateAPIResponse, error) {
	var resp media.AlibabaTjbPictureFolderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
