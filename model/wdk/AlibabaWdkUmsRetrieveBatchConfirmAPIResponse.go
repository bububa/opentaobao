package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkUmsRetrieveBatchConfirmAPIResponse 批量消息确认 API返回值
// alibaba.wdk.ums.retrieve.batch.confirm
//
// 批量消息确认
type AlibabaWdkUmsRetrieveBatchConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaWdkUmsRetrieveBatchConfirmAPIResponseModel
}

// AlibabaWdkUmsRetrieveBatchConfirmAPIResponseModel is 批量消息确认 成功返回结果
type AlibabaWdkUmsRetrieveBatchConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_retrieve_batch_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`
}
