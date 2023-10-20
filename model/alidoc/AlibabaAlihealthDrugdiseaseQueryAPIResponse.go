package alidoc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugdiseasequeryAPIResponse 药品诊断查询 API返回值
// alibaba.alihealth.drugdisease.query
//
// 药品诊断查询
type AlibabaalihealthdrugdiseasequeryAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugdiseasequeryAPIResponseModel
}

// AlibabaalihealthdrugdiseasequeryAPIResponseModel is 药品诊断查询 成功返回结果
type AlibabaalihealthdrugdiseasequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugdisease_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Data []DrugDiseaseMappingDto `json:"data,omitempty" xml:"data>drug_disease_mapping_dto,omitempty"`
	// 结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
