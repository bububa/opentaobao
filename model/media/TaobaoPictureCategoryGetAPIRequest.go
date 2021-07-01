package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureCategoryGetAPIRequest
获取图片分类信息 API请求
taobao.picture.category.get

获取图片分类信息 */
type TaobaoPictureCategoryGetAPIRequest struct {
	model.Params
	// 1
	_type string
	// 图片分类ID
	_pictureCategoryId int64
	// 图片分类名，不支持模糊查询
	_pictureCategoryName string
	// 取二级分类时设置为对应父分类id<br/>取一级分类时父分类id设为0<br/>取全部分类的时候不设或设为-1
	_parentId int64
	// 图片分类的修改时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片分类。
	_modifiedTime string
}

// New
