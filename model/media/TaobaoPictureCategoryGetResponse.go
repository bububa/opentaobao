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
    TaobaoPictureCategoryGetResponse
}

type TaobaoPictureCategoryGetResponse struct {
    XMLName xml.Name `xml:"picture_category_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 图片分类
    
    PictureCategories   []PictureCategory `json:"picture_categories,omitempty" xml:"picture_categories>picture_category,omitempty"`
    
    
}
