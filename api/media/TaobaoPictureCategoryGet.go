package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Taobaopicturecategoryget 获取图片分类信息
// taobao.picture.category.get
//
// 获取图片分类信息
func Taobaopicturecategoryget(clt *core.SDKClient, req *media.TaobaopicturecategorygetAPIRequest, session string) (*media.TaobaopicturecategorygetAPIResponse, error) {
	var resp media.TaobaopicturecategorygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
