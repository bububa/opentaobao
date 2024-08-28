package media

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPictureUserinfoGet 查询图片空间用户的信息
// taobao.picture.userinfo.get
//
// 查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量
func TaobaoPictureUserinfoGet(ctx context.Context, clt *core.SDKClient, req *media.TaobaoPictureUserinfoGetAPIRequest, resp *media.TaobaoPictureUserinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
