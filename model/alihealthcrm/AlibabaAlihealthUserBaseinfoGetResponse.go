package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户基础信息 APIResponse
alibaba.alihealth.user.baseinfo.get

获取用户基础信息
*/
type AlibabaAlihealthUserBaseinfoGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthUserBaseinfoGetResponse
}

type AlibabaAlihealthUserBaseinfoGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_user_baseinfo_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
