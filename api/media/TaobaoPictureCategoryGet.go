package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPictureCategoryGet 获取图片分类信息
// taobao.picture.category.get
//
// 获取图片分类信息
func TaobaoPictureCategoryGet(clt *core.SDKClient, req *media.TaobaoPictureCategoryGetAPIRequest, resp *media.TaobaoPictureCategoryGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
