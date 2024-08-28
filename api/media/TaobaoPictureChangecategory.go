package media

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPictureChangecategory 修改图片的分类
// taobao.picture.changecategory
//
// 把批量的图片移动到某个分类下
func TaobaoPictureChangecategory(ctx context.Context, clt *core.SDKClient, req *media.TaobaoPictureChangecategoryAPIRequest, resp *media.TaobaoPictureChangecategoryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
