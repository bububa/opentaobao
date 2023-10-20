package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPictureGet 获取图片信息
// taobao.picture.get
//
// 获取图片信息
func TaobaoPictureGet(clt *core.SDKClient, req *media.TaobaoPictureGetAPIRequest, session string) (*media.TaobaoPictureGetAPIResponse, error) {
	var resp media.TaobaoPictureGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
