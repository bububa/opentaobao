package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

/* TaobaoPictureCategoryUpdate
更新图片分类
taobao.picture.category.update

更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。 */
func TaobaoPictureCategoryUpdate(clt *core.SDKClient, req *media.TaobaoPictureCategoryUpdateAPIRequest, session string) (*media.TaobaoPictureCategoryUpdateAPIResponse, error) {
	var resp media.TaobaoPictureCategoryUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
