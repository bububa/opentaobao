package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowGetbyverifycodeAPIResponse 处方外流药店通过核销码获取处方 API返回值
// alibaba.alihealth.outflow.getbyverifycode
//
// 阿里健康对合作药店提供通过核销码查看处方的功能
type AlibabaAlihealthOutflowGetbyverifycodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthOutflowGetbyverifycodeAPIResponseModel
}

// AlibabaAlihealthOutflowGetbyverifycodeAPIResponseModel is 处方外流药店通过核销码获取处方 成功返回结果
type AlibabaAlihealthOutflowGetbyverifycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_getbyverifycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
