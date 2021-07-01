package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureCategoryAddAPIRequest
新增图片分类信息 API请求
taobao.picture.category.add

同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符 */
type TaobaoPictureCategoryAddAPIRequest struct {
	model.Params
	// 图片分类名称，最大长度20字符，中文字符算2个字符，不能为空
	_pictureCategoryName string
	// 图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id
	_parentId int64
}

// New
