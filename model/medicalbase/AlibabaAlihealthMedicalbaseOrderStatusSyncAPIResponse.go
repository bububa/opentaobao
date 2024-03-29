package medicalbase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse 号源直连订单状态同步接口 API返回值
// alibaba.alihealth.medicalbase.order.status.sync
//
// 互联网医院isv批量通过接口批量导入
type AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponseModel is 号源直连订单状态同步接口 成功返回结果
type AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_order_status_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse
func GetAlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse() *AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse {
	return poolAlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse.Get().(*AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse 将 AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse(v *AlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalbaseOrderStatusSyncAPIResponse.Put(v)
}
