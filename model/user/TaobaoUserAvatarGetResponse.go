package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝用户头像查询 APIResponse
taobao.user.avatar.get

根据混淆nick查询用户头像
*/
type TaobaoUserAvatarGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUserAvatarGetResponse `json:"taobao_user_avatar_get_response,omitempty"`
}

type TaobaoUserAvatarGetResponse struct {

    // 用户头像地址
    Avatar   string `json:"avatar,omitempty"`

}
