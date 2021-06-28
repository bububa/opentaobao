package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询图片空间用户的信息 APIResponse
taobao.picture.userinfo.get

查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量
*/
type TaobaoPictureUserinfoGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"picture_userinfo_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 用户使用图片空间的信息
    
    UserInfo   *UserInfo `json:"user_info,omitempty" xml:"