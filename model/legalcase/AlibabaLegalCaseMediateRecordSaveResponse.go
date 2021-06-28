package legalcase

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增调解结果 APIResponse
alibaba.legal.case.mediate.record.save

增加调解沟通记录
*/
type AlibabaLegalCaseMediateRecordSaveAPIResponse struct {
    model.CommonResponse
    AlibabaLegalCaseMediateRecordSaveResponse
}

type AlibabaLegalCaseMediateRecordSaveResponse struct {
    XMLName xml.Name `xml:"alibaba_legal_case_mediate_record_save_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
