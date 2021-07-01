package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureCategoryUpdateAPIRequest
更新图片分类 API请求
taobao.picture.category.update

更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。 */
type TaobaoPictureCategoryUpdateAPIRequest struct {
	model.Params
	// 要更新的图片分类的id
	_categoryId int64
	// 图片分类的新名字，最大长度20字符，不能为空
	_categoryName string
	// 图片分类的新父分类id
	_parentId int64
}

// New
