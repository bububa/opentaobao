package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取用户公钥 APIResponse
account.aliyuncs.com.GetPubKey.2013-07-01

根据用户的appkey查询用户的pubkey
*/
type AccountAliyuncsComGetPubKey2013-07-01APIResponse struct {
    model.CommonResponse
    // Response *AccountAliyuncsComGetPubKey2013-07-01Response `json:"account_aliyuncs_com_GetPubKey_2013-07-01_response,omitempty"` 
    AccountAliyuncsComGetPubKey2013-07-01Response
}

/* model for simplify = false
type AccountAliyuncsComGetPubKey2013-07-01Response struct {

    // 用户的公钥
    
    PubKey   string `json:"PubKey,omitempty"`
    

    // appkey的OwnerId
    
    OwnerId   string `json:"OwnerId,omitempty"`
    

    // 用户appkey的类型
    
    KeyType   string `json:"KeyType,omitempty"`
    

    // 创建key的时间
    
    CreateTime   string `json:"CreateTime,omitempty"`
    

    // RequestId
    
    RequestId   string `json:"RequestId,omitempty"`
    

    // Message
    
    Message   string `json:"Message,omitempty"`
    

    // 返回code
    
    Code   string `json:"Code,omitempty"`
    

    // 用户的appkey
    
    AppKey   string `json:"AppKey,omitempty"`
    

}
*/

type AccountAliyuncsComGetPubKey2013-07-01Response struct {

    // 用户的公钥
    PubKey   string `json:"PubKey,omitempty"`

    // appkey的OwnerId
    OwnerId   string `json:"OwnerId,omitempty"`

    // 用户appkey的类型
    KeyType   string `json:"KeyType,omitempty"`

    // 创建key的时间
    CreateTime   string `json:"CreateTime,omitempty"`

    // RequestId
    RequestId   string `json:"RequestId,omitempty"`

    // Message
    Message   string `json:"Message,omitempty"`

    // 返回code
    Code   string `json:"Code,omitempty"`

    // 用户的appkey
    AppKey   string `json:"AppKey,omitempty"`

}
