package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

/* TaobaoPictureCategoryGet
获取图片分类信息
taobao.picture.category.get

获取图片分类信息 */
func TaobaoPictureCategoryGet(clt *core.SDKClient, req *media.TaobaoPictureCategoryGetAPIRequest, session string) (*media.TaobaoPictureCategoryGetAPIResponse, error) {
	var resp media.TaobaoPictureCategoryGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
