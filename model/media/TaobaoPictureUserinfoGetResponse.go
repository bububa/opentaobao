package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询图片空间用户的信息 APIResponse
taobao.picture.userinfo.get

查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量
*/
type TaobaoPictureUserinfoGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPictureUserinfoGetResponse `json:"picture_userinfo_get_response,omitempty"` 
    TaobaoPictureUserinfoGetResponse
}

/* model for simplify = false
type TaobaoPictureUserinfoGetResponse struct {

    // 用户使用图片空间的信息
    
    UserInfo  *struct {
        UserInfo  *UserInfo `json:"user_info,omitempty"`
    } `json:"user_info,omitempty"`
    

}
*/

type TaobaoPictureUserinfoGetResponse struct {

    // 用户使用图片空间的信息
    UserInfo   *UserInfo `json:"user_info,omitempty"`

}
