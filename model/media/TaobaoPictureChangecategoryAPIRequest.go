package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureChangecategoryAPIRequest
修改图片的分类 API请求
taobao.picture.changecategory

把批量的图片移动到某个分类下 */
type TaobaoPictureChangecategoryAPIRequest struct {
	model.Params
	// 要移动的图片的id
	_pictureIds []int64
	// 目标分类的id
	_pictureCategoryId int64
}

// New
