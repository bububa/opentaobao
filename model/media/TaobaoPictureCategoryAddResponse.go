package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增图片分类信息 APIResponse
taobao.picture.category.add

同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符
*/
type TaobaoPictureCategoryAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPictureCategoryAddResponse `json:"picture_category_add_response,omitempty"` 
    TaobaoPictureCategoryAddResponse
}

/* model for simplify = false
type TaobaoPictureCategoryAddResponse struct {

    // 图片分类信息
    
    PictureCategory  *struct {
        PictureCategory  *PictureCategory `json:"picture_category,omitempty"`
    } `json:"picture_category,omitempty"`
    

}
*/

type TaobaoPictureCategoryAddResponse struct {

    // 图片分类信息
    PictureCategory   *PictureCategory `json:"picture_category,omitempty"`

}
