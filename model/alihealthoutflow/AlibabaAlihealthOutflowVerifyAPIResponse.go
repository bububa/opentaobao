package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowVerifyAPIResponse 处方外流药店通过核销码核销处方 API返回值
// alibaba.alihealth.outflow.verify
//
// 处方外流药店通过核销码核销处方
type AlibabaAlihealthOutflowVerifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthOutflowVerifyAPIResponseModel
}

// AlibabaAlihealthOutflowVerifyAPIResponseModel is 处方外流药店通过核销码核销处方 成功返回结果
type AlibabaAlihealthOutflowVerifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
