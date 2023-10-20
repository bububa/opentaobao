package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Taobaopictureget 获取图片信息
// taobao.picture.get
//
// 获取图片信息
func Taobaopictureget(clt *core.SDKClient, req *media.TaobaopicturegetAPIRequest, session string) (*media.TaobaopicturegetAPIResponse, error) {
	var resp media.TaobaopicturegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
