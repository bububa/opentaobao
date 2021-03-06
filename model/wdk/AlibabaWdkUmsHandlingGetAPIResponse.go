package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsHandlingGetAPIResponse 加工单-回流单（新接口） API返回值
// alibaba.wdk.ums.handling.get
//
// 加工单-回流单（新接口）
type AlibabaWdkUmsHandlingGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsHandlingGetAPIResponseModel
}

// AlibabaWdkUmsHandlingGetAPIResponseModel is 加工单-回流单（新接口） 成功返回结果
type AlibabaWdkUmsHandlingGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_handling_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
