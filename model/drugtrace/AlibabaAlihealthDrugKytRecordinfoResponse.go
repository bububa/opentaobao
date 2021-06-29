package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
快易通健康检查 API返回值 
alibaba.alihealth.drug.kyt.recordinfo

快易通健康检查
*/
type AlibabaAlihealthDrugKytRecordinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytRecordinfoResponse
}

// 快易通健康检查 成功返回结果
type AlibabaAlihealthDrugKytRecordinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_recordinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 对象
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
    // 返回码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 返回值
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 是否响应成功
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
