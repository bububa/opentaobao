package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改图片的分类 API请求
taobao.picture.changecategory

把批量的图片移动到某个分类下
*/
type TaobaoPictureChangecategoryRequest struct {
    model.Params
    // 要移动的图片的id
    pictureIds   []int64
    // 目标分类的id
    pictureCategoryId   int64
}

// 初始化TaobaoPictureChangecategoryRequest对象
func NewTaobaoPictureChangecategoryRequest() *TaobaoPictureChangecategoryRequest{
    return &TaobaoPictureChangecategoryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPictureChangecategoryRequest) GetApiMethodName() string {
    return "taobao.picture.changecategory"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPictureChangecategoryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PictureIds Setter
// 要移动的图片的id
func (r *TaobaoPictureChangecategoryRequest) SetPictureIds(pictureIds []int64) error {
    r.pictureIds = pictureIds
    r.Set("picture_ids", pictureIds)
    return nil
}

// PictureIds Getter
func (r TaobaoPictureChangecategoryRequest) GetPictureIds() []int64 {
    return r.pictureIds
}
// PictureCategoryId Setter
// 目标分类的id
func (r *TaobaoPictureChangecategoryRequest) SetPictureCategoryId(pictureCategoryId int64) error {
    r.pictureCategoryId = pictureCategoryId
    r.Set("picture_category_id", pictureCategoryId)
    return nil
}

// PictureCategoryId Getter
func (r TaobaoPictureChangecategoryRequest) GetPictureCategoryId() int64 {
    return r.pictureCategoryId
}
