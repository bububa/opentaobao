package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowDrugSupervisionQueryAPIResponse 监管平台药品查询 API返回值
// alibaba.alihealth.outflow.drug.supervision.query
//
// 获取监管平台药品数据
type AlibabaAlihealthOutflowDrugSupervisionQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthOutflowDrugSupervisionQueryAPIResponseModel
}

// AlibabaAlihealthOutflowDrugSupervisionQueryAPIResponseModel is 监管平台药品查询 成功返回结果
type AlibabaAlihealthOutflowDrugSupervisionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_drug_supervision_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
