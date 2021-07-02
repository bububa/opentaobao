package legalcase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseMediateRecordSaveAPIResponse 新增调解结果 API返回值
// alibaba.legal.case.mediate.record.save
//
// 增加调解沟通记录
type AlibabaLegalCaseMediateRecordSaveAPIResponse struct {
	model.CommonResponse
	AlibabaLegalCaseMediateRecordSaveAPIResponseModel
}

// AlibabaLegalCaseMediateRecordSaveAPIResponseModel is 新增调解结果 成功返回结果
type AlibabaLegalCaseMediateRecordSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_mediate_record_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
