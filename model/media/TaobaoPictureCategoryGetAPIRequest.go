package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureCategoryGetAPIRequest 获取图片分类信息 API请求
// taobao.picture.category.get
//
// 获取图片分类信息
type TaobaoPictureCategoryGetAPIRequest struct {
	model.Params
	// 图片分类名，不支持模糊查询
	_pictureCategoryName string
	// 1
	_type string
	// 图片分类的修改时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片分类。
	_modifiedTime string
	// 图片分类ID
	_pictureCategoryId int64
	// 取二级分类时设置为对应父分类id<br/>取一级分类时父分类id设为0<br/>取全部分类的时候不设或设为-1
	_parentId int64
}

// NewTaobaoPictureCategoryGetRequest 初始化TaobaoPictureCategoryGetAPIRequest对象
func NewTaobaoPictureCategoryGetRequest() *TaobaoPictureCategoryGetAPIRequest {
	return &TaobaoPictureCategoryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureCategoryGetAPIRequest) GetApiMethodName() string {
	return "taobao.picture.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureCategoryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPictureCategoryName is PictureCategoryName Setter
// 图片分类名，不支持模糊查询
func (r *TaobaoPictureCategoryGetAPIRequest) SetPictureCategoryName(_pictureCategoryName string) error {
	r._pictureCategoryName = _pictureCategoryName
	r.Set("picture_category_name", _pictureCategoryName)
	return nil
}

// GetPictureCategoryName PictureCategoryName Getter
func (r TaobaoPictureCategoryGetAPIRequest) GetPictureCategoryName() string {
	return r._pictureCategoryName
}

// SetType is Type Setter
// 1
func (r *TaobaoPictureCategoryGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoPictureCategoryGetAPIRequest) GetType() string {
	return r._type
}

// SetModifiedTime is ModifiedTime Setter
// 图片分类的修改时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片分类。
func (r *TaobaoPictureCategoryGetAPIRequest) SetModifiedTime(_modifiedTime string) error {
	r._modifiedTime = _modifiedTime
	r.Set("modified_time", _modifiedTime)
	return nil
}

// GetModifiedTime ModifiedTime Getter
func (r TaobaoPictureCategoryGetAPIRequest) GetModifiedTime() string {
	return r._modifiedTime
}

// SetPictureCategoryId is PictureCategoryId Setter
// 图片分类ID
func (r *TaobaoPictureCategoryGetAPIRequest) SetPictureCategoryId(_pictureCategoryId int64) error {
	r._pictureCategoryId = _pictureCategoryId
	r.Set("picture_category_id", _pictureCategoryId)
	return nil
}

// GetPictureCategoryId PictureCategoryId Getter
func (r TaobaoPictureCategoryGetAPIRequest) GetPictureCategoryId() int64 {
	return r._pictureCategoryId
}

// SetParentId is ParentId Setter
// 取二级分类时设置为对应父分类id&lt;br/&gt;取一级分类时父分类id设为0&lt;br/&gt;取全部分类的时候不设或设为-1
func (r *TaobaoPictureCategoryGetAPIRequest) SetParentId(_parentId int64) error {
	r._parentId = _parentId
	r.Set("parent_id", _parentId)
	return nil
}

// GetParentId ParentId Getter
func (r TaobaoPictureCategoryGetAPIRequest) GetParentId() int64 {
	return r._parentId
}
