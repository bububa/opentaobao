package alihealth2

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponseModel is 第三方订单同步 成功返回结果
type AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_third_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *AlibabaAlihealthMedicalbaseThirdOrderSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse
func GetAlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse() *AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse {
	return poolAlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse.Get().(*AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse 将 AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse(v *AlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseThirdOrderSyncAPIResponse.Put(v)
}
