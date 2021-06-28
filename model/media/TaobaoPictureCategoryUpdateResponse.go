package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新图片分类 APIResponse
taobao.picture.category.update

更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。
*/
type TaobaoPictureCategoryUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoPictureCategoryUpdateResponse
}

type TaobaoPictureCategoryUpdateResponse struct {
    XMLName xml.Name `xml:"picture_category_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 更新图片分类是否成功
    
    Done   bool `json:"done,omitempty" xml:"done,omitempty"`

    
}
