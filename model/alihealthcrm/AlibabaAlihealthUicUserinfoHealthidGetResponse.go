package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取健康id APIResponse
alibaba.alihealth.uic.userinfo.healthid.get

根据支付宝用户ID获取用户健康ID
*/
type AlibabaAlihealthUicUserinfoHealthidGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthUicUserinfoHealthidGetResponse
}

type AlibabaAlihealthUicUserinfoHealthidGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_uic_userinfo_healthid_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
