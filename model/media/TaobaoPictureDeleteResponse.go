package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除图片空间图片 APIResponse
taobao.picture.delete

删除图片空间图片
*/
type TaobaoPictureDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPictureDeleteResponse `json:"picture_delete_response,omitempty"` 
    TaobaoPictureDeleteResponse
}

/* model for simplify = false
type TaobaoPictureDeleteResponse struct {

    // 是否删除
    
    Success   string `json:"success,omitempty"`
    

}
*/

type TaobaoPictureDeleteResponse struct {

    // 是否删除
    Success   string `json:"success,omitempty"`

}
