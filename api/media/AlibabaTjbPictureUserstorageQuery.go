package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Alibabatjbpictureuserstoragequery 淘特图片空间用户容量查询
// alibaba.tjb.picture.userstorage.query
//
// 淘特图片空间用户容量查询
func Alibabatjbpictureuserstoragequery(clt *core.SDKClient, req *media.AlibabatjbpictureuserstoragequeryAPIRequest, session string) (*media.AlibabatjbpictureuserstoragequeryAPIResponse, error) {
	var resp media.AlibabatjbpictureuserstoragequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
