package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取图片分类信息 APIResponse
taobao.picture.category.get

获取图片分类信息
*/
type TaobaoPictureCategoryGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPictureCategoryGetResponse `json:"taobao_picture_category_get_response,omitempty"`
}

type TaobaoPictureCategoryGetResponse struct {

    // 图片分类
    PictureCategories   []PictureCategory `json:"picture_categories,omitempty"`

}
