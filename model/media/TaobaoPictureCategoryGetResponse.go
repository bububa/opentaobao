package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取图片分类信息 APIResponse
taobao.picture.category.get

获取图片分类信息
*/
type TaobaoPictureCategoryGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"picture_category_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 图片分类
    
    PictureCategories   []PictureCategory `json:"picture_categories,omitempty" xml:"