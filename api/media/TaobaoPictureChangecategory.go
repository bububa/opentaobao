package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Taobaopicturechangecategory 修改图片的分类
// taobao.picture.changecategory
//
// 把批量的图片移动到某个分类下
func Taobaopicturechangecategory(clt *core.SDKClient, req *media.TaobaopicturechangecategoryAPIRequest, session string) (*media.TaobaopicturechangecategoryAPIResponse, error) {
	var resp media.TaobaopicturechangecategoryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
