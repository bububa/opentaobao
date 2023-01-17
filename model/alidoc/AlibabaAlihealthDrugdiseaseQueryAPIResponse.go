package alidoc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugdiseaseQueryAPIResponse 药品诊断查询 API返回值
// alibaba.alihealth.drugdisease.query
//
// 药品诊断查询
type AlibabaAlihealthDrugdiseaseQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugdiseaseQueryAPIResponseModel
}

// AlibabaAlihealthDrugdiseaseQueryAPIResponseModel is 药品诊断查询 成功返回结果
type AlibabaAlihealthDrugdiseaseQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugdisease_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Data []DrugDiseaseMappingDto `json:"data,omitempty" xml:"data>drug_disease_mapping_dto,omitempty"`
	// 结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
