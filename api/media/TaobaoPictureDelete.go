package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPictureDelete 删除图片空间图片
// taobao.picture.delete
//
// 删除图片空间图片
func TaobaoPictureDelete(clt *core.SDKClient, req *media.TaobaoPictureDeleteAPIRequest, resp *media.TaobaoPictureDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
