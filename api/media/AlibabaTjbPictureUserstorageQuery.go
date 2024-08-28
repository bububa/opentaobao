package media

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// AlibabaTjbPictureUserstorageQuery 淘特图片空间用户容量查询
// alibaba.tjb.picture.userstorage.query
//
// 淘特图片空间用户容量查询
func AlibabaTjbPictureUserstorageQuery(ctx context.Context, clt *core.SDKClient, req *media.AlibabaTjbPictureUserstorageQueryAPIRequest, resp *media.AlibabaTjbPictureUserstorageQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
