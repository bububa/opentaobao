package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户开放信息获取 APIResponse
taobao.miniapp.userInfo.get

获取用户的 openId，snsNick（如果用户设置过的话），和加密头像链接
*/
type TaobaoMiniappUserInfoGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"miniapp_userInfo_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoMiniappUserInfoGetResult `json:"result,omitempty" xml:"