package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改图片名字 APIResponse
taobao.picture.update

修改指定图片的图片名
*/
type TaobaoPictureUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPictureUpdateResponse `json:"picture_update_response,omitempty"` 
    TaobaoPictureUpdateResponse
}

/* model for simplify = false
type TaobaoPictureUpdateResponse struct {

    // 更新是否成功
    
    Done   bool `json:"done,omitempty"`
    

}
*/

type TaobaoPictureUpdateResponse struct {

    // 更新是否成功
    Done   bool `json:"done,omitempty"`

}
