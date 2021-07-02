package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsRetrieveConfirmAPIResponse 回流单－外部对已拉取到的UMS单据进行确认 API返回值
// alibaba.wdk.ums.retrieve.confirm
//
// 回流单－外部对已拉取到的UMS单据进行确认
type AlibabaWdkUmsRetrieveConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsRetrieveConfirmAPIResponseModel
}

// AlibabaWdkUmsRetrieveConfirmAPIResponseModel is 回流单－外部对已拉取到的UMS单据进行确认 成功返回结果
type AlibabaWdkUmsRetrieveConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_retrieve_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
