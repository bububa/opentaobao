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
    AccountAliyuncsComGetPubKey2013-07-01Response
}

type AccountAliyuncsComGetPubKey2013-07-01Response struct {
    XMLName xml.Name `xml:"account_aliyuncs_com_GetPubKey_2013-07-01_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 用户的公钥
    
    PubKey   string `json:"PubKey,omitempty" xml:"PubKey,omitempty"`

    
    // appkey的OwnerId
    
    OwnerId   string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`

    
    // 用户appkey的类型
    
    KeyType   string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`

    
    // 创建key的时间
    
    CreateTime   string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`

    
    // RequestId
    
    RequestId   string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`

    
    // Message
    
    Message   string `json:"Message,omitempty" xml:"Message,omitempty"`

    
    // 返回code
    
    Code   string `json:"Code,omitempty" xml:"Code,omitempty"`

    
    // 用户的appkey
    
    AppKey   string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`

    
}
