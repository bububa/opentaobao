package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaTjbPictureUserstorageQuery 淘特图片空间用户容量查询
// alibaba.tjb.picture.userstorage.query
//
// 淘特图片空间用户容量查询
func AlibabaTjbPictureUserstorageQuery(clt *core.SDKClient, req *media.AlibabaTjbPictureUserstorageQueryAPIRequest, session string) (*media.AlibabaTjbPictureUserstorageQueryAPIResponse, error) {
	var resp media.AlibabaTjbPictureUserstorageQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
