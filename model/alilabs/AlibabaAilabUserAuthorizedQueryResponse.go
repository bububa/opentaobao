package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询授权状态接口 APIResponse
alibaba.ailab.user.authorized.query

查询三方用户授权状态
*/
type AlibabaAilabUserAuthorizedQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAilabUserAuthorizedQueryResponse
}

type AlibabaAilabUserAuthorizedQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ailab_user_authorized_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 200 成功，其他失败
    
    StatusCode   int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
    // 是否已授权
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
