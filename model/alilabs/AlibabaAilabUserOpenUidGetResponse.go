package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
access token 获取精灵用户 id APIResponse
alibaba.ailab.user.open.uid.get

access token 获取精灵用户 id
*/
type AlibabaAilabUserOpenUidGetAPIResponse struct {
    model.CommonResponse
    AlibabaAilabUserOpenUidGetResponse
}

type AlibabaAilabUserOpenUidGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ailab_user_open_uid_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 详细信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // user id
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
    // 状态码，200 成功，其他失败
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
}
