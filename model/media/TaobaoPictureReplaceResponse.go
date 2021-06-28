package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
替换图片 APIResponse
taobao.picture.replace

替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。
*/
type TaobaoPictureReplaceAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPictureReplaceResponse `json:"picture_replace_response,omitempty"` 
    TaobaoPictureReplaceResponse
}

/* model for simplify = false
type TaobaoPictureReplaceResponse struct {

    // 图片替换是否成功
    
    Done   bool `json:"done,omitempty"`
    

}
*/

type TaobaoPictureReplaceResponse struct {

    // 图片替换是否成功
    Done   bool `json:"done,omitempty"`

}
