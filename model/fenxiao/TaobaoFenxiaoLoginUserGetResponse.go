package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取分销用户登录信息 API返回值 
taobao.fenxiao.login.user.get

获取用户登录信息
*/
type TaobaoFenxiaoLoginUserGetAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoLoginUserGetResponse
}

// 获取分销用户登录信息 成功返回结果
type TaobaoFenxiaoLoginUserGetResponse struct {
    XMLName xml.Name `xml:"fenxiao_login_user_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 登录用户信息
    LoginUser   *LoginUser `json:"login_user,omitempty" xml:"login_user,omitempty"`
}
