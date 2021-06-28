package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询买家信息API APIResponse
taobao.user.buyer.get

查询买家信息API，只能买家类应用调用。
*/
type TaobaoUserBuyerGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"user_buyer_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 只返回nick,avatar参数
    
    User   *User `json:"user,omitempty" xml:"