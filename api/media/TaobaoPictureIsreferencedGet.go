package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Taobaopictureisreferencedget 图片是否被引用
// taobao.picture.isreferenced.get
//
// 查询图片是否被引用，被引用返回true，未被引用返回false
func Taobaopictureisreferencedget(clt *core.SDKClient, req *media.TaobaopictureisreferencedgetAPIRequest, session string) (*media.TaobaopictureisreferencedgetAPIResponse, error) {
	var resp media.TaobaopictureisreferencedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
