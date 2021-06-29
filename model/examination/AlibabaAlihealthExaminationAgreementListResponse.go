package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
isv协议获取 API返回值 
alibaba.alihealth.examination.agreement.list

isv协议获取
*/
type AlibabaAlihealthExaminationAgreementListAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationAgreementListResponse
}

// isv协议获取 成功返回结果
type AlibabaAlihealthExaminationAgreementListResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_agreement_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果编码
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
    // 返回结果描述
    ResponseMessage   string `json:"response_message,omitempty" xml:"response_message,omitempty"`
    // 返回的json格式数据
    Agreement   *Agreement `json:"agreement,omitempty" xml:"agreement,omitempty"`
}
