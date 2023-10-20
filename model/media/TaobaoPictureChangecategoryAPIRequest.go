package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureChangecategoryAPIRequest 修改图片的分类 API请求
// taobao.picture.changecategory
//
// 把批量的图片移动到某个分类下
type TaobaoPictureChangecategoryAPIRequest struct {
	model.Params
	// 要移动的图片的id
	_pictureIds []string
	// 目标分类的id
	_pictureCategoryId int64
}

// NewTaobaoPictureChangecategoryRequest 初始化TaobaoPictureChangecategoryAPIRequest对象
func NewTaobaoPictureChangecategoryRequest() *TaobaoPictureChangecategoryAPIRequest {
	return &TaobaoPictureChangecategoryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureChangecategoryAPIRequest) GetApiMethodName() string {
	return "taobao.picture.changecategory"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureChangecategoryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPictureChangecategoryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPictureIds is PictureIds Setter
// 要移动的图片的id
func (r *TaobaoPictureChangecategoryAPIRequest) SetPictureIds(_pictureIds []string) error {
	r._pictureIds = _pictureIds
	r.Set("picture_ids", _pictureIds)
	return nil
}

// GetPictureIds PictureIds Getter
func (r TaobaoPictureChangecategoryAPIRequest) GetPictureIds() []string {
	return r._pictureIds
}

// SetPictureCategoryId is PictureCategoryId Setter
// 目标分类的id
func (r *TaobaoPictureChangecategoryAPIRequest) SetPictureCategoryId(_pictureCategoryId int64) error {
	r._pictureCategoryId = _pictureCategoryId
	r.Set("picture_category_id", _pictureCategoryId)
	return nil
}

// GetPictureCategoryId PictureCategoryId Getter
func (r TaobaoPictureChangecategoryAPIRequest) GetPictureCategoryId() int64 {
	return r._pictureCategoryId
}
