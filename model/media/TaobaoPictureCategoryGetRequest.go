package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取图片分类信息 API请求
taobao.picture.category.get

获取图片分类信息
*/
type TaobaoPictureCategoryGetRequest struct {
    model.Params
    // 1
    type   string
    // 图片分类ID
    pictureCategoryId   int64
    // 图片分类名，不支持模糊查询
    pictureCategoryName   string
    // 取二级分类时设置为对应父分类id<br/>取一级分类时父分类id设为0<br/>取全部分类的时候不设或设为-1
    parentId   int64
    // 图片分类的修改时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片分类。
    modifiedTime   string
}

// 初始化TaobaoPictureCategoryGetRequest对象
func NewTaobaoPictureCategoryGetRequest() *TaobaoPictureCategoryGetRequest{
    return &TaobaoPictureCategoryGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPictureCategoryGetRequest) GetApiMethodName() string {
    return "taobao.picture.category.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPictureCategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 1
func (r *TaobaoPictureCategoryGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoPictureCategoryGetRequest) GetType() string {
    return r.type
}
// PictureCategoryId Setter
// 图片分类ID
func (r *TaobaoPictureCategoryGetRequest) SetPictureCategoryId(pictureCategoryId int64) error {
    r.pictureCategoryId = pictureCategoryId
    r.Set("picture_category_id", pictureCategoryId)
    return nil
}

// PictureCategoryId Getter
func (r TaobaoPictureCategoryGetRequest) GetPictureCategoryId() int64 {
    return r.pictureCategoryId
}
// PictureCategoryName Setter
// 图片分类名，不支持模糊查询
func (r *TaobaoPictureCategoryGetRequest) SetPictureCategoryName(pictureCategoryName string) error {
    r.pictureCategoryName = pictureCategoryName
    r.Set("picture_category_name", pictureCategoryName)
    return nil
}

// PictureCategoryName Getter
func (r TaobaoPictureCategoryGetRequest) GetPictureCategoryName() string {
    return r.pictureCategoryName
}
// ParentId Setter
// 取二级分类时设置为对应父分类id<br/>取一级分类时父分类id设为0<br/>取全部分类的时候不设或设为-1
func (r *TaobaoPictureCategoryGetRequest) SetParentId(parentId int64) error {
    r.parentId = parentId
    r.Set("parent_id", parentId)
    return nil
}

// ParentId Getter
func (r TaobaoPictureCategoryGetRequest) GetParentId() int64 {
    return r.parentId
}
// ModifiedTime Setter
// 图片分类的修改时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片分类。
func (r *TaobaoPictureCategoryGetRequest) SetModifiedTime(modifiedTime string) error {
    r.modifiedTime = modifiedTime
    r.Set("modified_time", modifiedTime)
    return nil
}

// ModifiedTime Getter
func (r TaobaoPictureCategoryGetRequest) GetModifiedTime() string {
    return r.modifiedTime
}
