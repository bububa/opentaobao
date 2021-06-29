package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发回帖子信息同步 APIResponse
alibaba.alihealth.pregnancy.posts.data

发回帖子信息同步
*/
type AlibabaAlihealthPregnancyPostsDataAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPregnancyPostsDataResponse
}

type AlibabaAlihealthPregnancyPostsDataResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_posts_data_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
}
