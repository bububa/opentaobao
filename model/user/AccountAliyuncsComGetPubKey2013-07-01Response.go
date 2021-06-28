package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户公钥 APIResponse
account.aliyuncs.com.GetPubKey.2013-07-01

根据用户的appkey查询用户的pubkey
*/
type AccountAliyuncsComGetPubKey2013-07-01APIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"account_aliyuncs_com_GetPubKey_2013-07-01_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 用户的公钥
    
    PubKey   string `json:"PubKey,omitempty" xml:"