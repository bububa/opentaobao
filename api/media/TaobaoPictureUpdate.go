package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPictureUpdate 修改图片名字
// taobao.picture.update
//
// 修改指定图片的图片名
func TaobaoPictureUpdate(clt *core.SDKClient, req *media.TaobaoPictureUpdateAPIRequest, resp *media.TaobaoPictureUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
