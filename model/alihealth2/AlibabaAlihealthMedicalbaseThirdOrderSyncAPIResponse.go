package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse 第三方订单同步 API返回值
// alibaba.alihealth.medicalbase.third.order.sync
//
// 第三方订单同步
type AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponseModel
}

// AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponseModel is 第三方订单同步 成功返回结果
type AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_third_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *AlibabaAlihealthMedicalbaseThirdOrderSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
