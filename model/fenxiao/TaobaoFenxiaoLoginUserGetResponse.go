package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取分销用户登录信息 APIResponse
taobao.fenxiao.login.user.get

获取用户登录信息
*/
type TaobaoFenxiaoLoginUserGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_login_user_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 登录用户信息
    
    LoginUser   *LoginUser `json:"login_user,omitempty" xml:"