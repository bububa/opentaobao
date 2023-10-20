package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse 门店无隐形消费签约 API返回值
// alibaba.alihealth.dental.store.invisible.consume.update
//
// 门店无隐形消费签约
type AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponseModel).Reset()
}

// AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponseModel is 门店无隐形消费签约 成功返回结果
type AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_store_invisible_consume_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse)
	},
}

// GetAlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse
func GetAlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse() *AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse {
	return poolAlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse.Get().(*AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse)
}

// ReleaseAlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse 将 AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse(v *AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse.Put(v)
}
