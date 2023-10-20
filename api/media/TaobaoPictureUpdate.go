package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Taobaopictureupdate 修改图片名字
// taobao.picture.update
//
// 修改指定图片的图片名
func Taobaopictureupdate(clt *core.SDKClient, req *media.TaobaopictureupdateAPIRequest, session string) (*media.TaobaopictureupdateAPIResponse, error) {
	var resp media.TaobaopictureupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
