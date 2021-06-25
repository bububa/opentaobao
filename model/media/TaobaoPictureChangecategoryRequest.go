package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改图片的分类 APIRequest
taobao.picture.changecategory

把批量的图片移动到某个分类下
*/
type TaobaoPictureChangecategoryRequest struct {
    model.Params

    // 要移动的图片的id
    pictureIds   []Number 

    // 目标分类的id
    pictureCategoryId   int64 

}

func NewTaobaoPictureChangecategoryRequest() *TaobaoPictureChangecategoryRequest{
    return &TaobaoPictureChangecategoryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPictureChangecategoryRequest) GetApiMethodName() string {
    return "taobao.picture.changecategory"
}

func (r TaobaoPictureChangecategoryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPictureChangecategoryRequest) SetPictureIds(pictureIds []Number) error {
    r.pictureIds = pictureIds
    r.Set("picture_ids", pictureIds)
    return nil
}

func (r TaobaoPictureChangecategoryRequest) GetPictureIds() []Number {
    return r.pictureIds
}

func (r *TaobaoPictureChangecategoryRequest) SetPictureCategoryId(pictureCategoryId int64) error {
    r.pictureCategoryId = pictureCategoryId
    r.Set("picture_category_id", pictureCategoryId)
    return nil
}

func (r TaobaoPictureChangecategoryRequest) GetPictureCategoryId() int64 {
    return r.pictureCategoryId
}

