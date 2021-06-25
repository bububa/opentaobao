package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取分销用户登录信息 APIResponse
taobao.fenxiao.login.user.get

获取用户登录信息
*/
type TaobaoFenxiaoLoginUserGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoLoginUserGetResponse `json:"taobao_fenxiao_login_user_get_response,omitempty"`
}

type TaobaoFenxiaoLoginUserGetResponse struct {

    // 登录用户信息
    LoginUser   *LoginUser `json:"login_user,omitempty"`

}
