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
    // Response *TaobaoPictureCategoryGetResponse `json:"picture_category_get_response,omitempty"` 
    TaobaoPictureCategoryGetResponse
}

/* model for simplify = false
type TaobaoPictureCategoryGetResponse struct {

    // 图片分类
    
    PictureCategories  struct {
        PictureCategory  []PictureCategory `json:"picture_category,omitempty"`
    } `json:"picture_categories,omitempty"`
    

}
*/

type TaobaoPictureCategoryGetResponse struct {

    // 图片分类
    PictureCategories   []PictureCategory `json:"picture_categories,omitempty"`

}
