package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
切换用户 API返回值 
alibaba.ailabs.tmallgenie.auth.switchuser

设备切换授权用户
*/
type AlibabaAilabsTmallgenieAuthSwitchuserAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsTmallgenieAuthSwitchuserResponse
}

// 切换用户 成功返回结果
type AlibabaAilabsTmallgenieAuthSwitchuserResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_switchuser_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *AlibabaAilabsTmallgenieAuthSwitchuserResult `json:"result,omitempty" xml:"result,omitempty"`
}
