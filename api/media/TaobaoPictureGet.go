package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPictureGet 获取图片信息
// taobao.picture.get
//
// 获取图片信息
func TaobaoPictureGet(clt *core.SDKClient, req *media.TaobaoPictureGetAPIRequest, resp *media.TaobaoPictureGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
