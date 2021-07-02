package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPictureIsreferencedGet 图片是否被引用
// taobao.picture.isreferenced.get
//
// 查询图片是否被引用，被引用返回true，未被引用返回false
func TaobaoPictureIsreferencedGet(clt *core.SDKClient, req *media.TaobaoPictureIsreferencedGetAPIRequest, session string) (*media.TaobaoPictureIsreferencedGetAPIResponse, error) {
	var resp media.TaobaoPictureIsreferencedGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
