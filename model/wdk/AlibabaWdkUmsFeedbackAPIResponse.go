package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsFeedbackAPIResponse
质量反馈（入库辅助）-ERP下发单 API返回值
alibaba.wdk.ums.feedback

质量反馈（入库辅助）-ERP下发单 */
type AlibabaWdkUmsFeedbackAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsFeedbackAPIResponseModel
}

// AlibabaWdkUmsFeedbackAPIResponseModel is 质量反馈（入库辅助）-ERP下发单 成功返回结果
type AlibabaWdkUmsFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
