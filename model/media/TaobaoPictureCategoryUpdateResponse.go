package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新图片分类 APIResponse
taobao.picture.category.update

更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。
*/
type TaobaoPictureCategoryUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPictureCategoryUpdateResponse `json:"taobao_picture_category_update_response,omitempty"`
}

type TaobaoPictureCategoryUpdateResponse struct {

    // 更新图片分类是否成功
    Done   bool `json:"done,omitempty"`

}
