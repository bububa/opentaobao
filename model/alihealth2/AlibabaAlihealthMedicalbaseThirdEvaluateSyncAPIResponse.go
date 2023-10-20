package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse 三方评论信息同步 API返回值
// alibaba.alihealth.medicalbase.third.evaluate.sync
//
// 三方评论信息同步
type AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponseModel
}

// AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponseModel is 三方评论信息同步 成功返回结果
type AlibabaAlihealthMedicalbaseThirdEvaluateSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_third_evaluate_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
