package media

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureCategoryAddAPIRequest 新增图片分类信息 API请求
// taobao.picture.category.add
//
// 同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符
type TaobaoPictureCategoryAddAPIRequest struct {
	model.Params
	// 图片分类名称，最大长度20字符，中文字符算2个字符，不能为空
	_pictureCategoryName string
	// 图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id
	_parentId int64
}

// NewTaobaoPictureCategoryAddRequest 初始化TaobaoPictureCategoryAddAPIRequest对象
func NewTaobaoPictureCategoryAddRequest() *TaobaoPictureCategoryAddAPIRequest {
	return &TaobaoPictureCategoryAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPictureCategoryAddAPIRequest) Reset() {
	r._pictureCategoryName = ""
	r._parentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPictureCategoryAddAPIRequest) GetApiMethodName() string {
	return "taobao.picture.category.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPictureCategoryAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPictureCategoryAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPictureCategoryName is PictureCategoryName Setter
// 图片分类名称，最大长度20字符，中文字符算2个字符，不能为空
func (r *TaobaoPictureCategoryAddAPIRequest) SetPictureCategoryName(_pictureCategoryName string) error {
	r._pictureCategoryName = _pictureCategoryName
	r.Set("picture_category_name", _pictureCategoryName)
	return nil
}

// GetPictureCategoryName PictureCategoryName Getter
func (r TaobaoPictureCategoryAddAPIRequest) GetPictureCategoryName() string {
	return r._pictureCategoryName
}

// SetParentId is ParentId Setter
// 图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id
func (r *TaobaoPictureCategoryAddAPIRequest) SetParentId(_parentId int64) error {
	r._parentId = _parentId
	r.Set("parent_id", _parentId)
	return nil
}

// GetParentId ParentId Getter
func (r TaobaoPictureCategoryAddAPIRequest) GetParentId() int64 {
	return r._parentId
}

var poolTaobaoPictureCategoryAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPictureCategoryAddRequest()
	},
}

// GetTaobaoPictureCategoryAddRequest 从 sync.Pool 获取 TaobaoPictureCategoryAddAPIRequest
func GetTaobaoPictureCategoryAddAPIRequest() *TaobaoPictureCategoryAddAPIRequest {
	return poolTaobaoPictureCategoryAddAPIRequest.Get().(*TaobaoPictureCategoryAddAPIRequest)
}

// ReleaseTaobaoPictureCategoryAddAPIRequest 将 TaobaoPictureCategoryAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoPictureCategoryAddAPIRequest(v *TaobaoPictureCategoryAddAPIRequest) {
	v.Reset()
	poolTaobaoPictureCategoryAddAPIRequest.Put(v)
}
