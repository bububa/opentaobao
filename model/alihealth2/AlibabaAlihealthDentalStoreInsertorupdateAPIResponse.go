package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalStoreInsertorupdateAPIResponse ISV新增/修改口腔门店 API返回值
// alibaba.alihealth.dental.store.insertorupdate
//
// ISV新增/修改口腔门店
type AlibabaAlihealthDentalStoreInsertorupdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDentalStoreInsertorupdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalStoreInsertorupdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDentalStoreInsertorupdateAPIResponseModel).Reset()
}

// AlibabaAlihealthDentalStoreInsertorupdateAPIResponseModel is ISV新增/修改口腔门店 成功返回结果
type AlibabaAlihealthDentalStoreInsertorupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_store_insertorupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaAlihealthDentalStoreInsertorupdateMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDentalStoreInsertorupdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDentalStoreInsertorupdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalStoreInsertorupdateAPIResponse)
	},
}

// GetAlibabaAlihealthDentalStoreInsertorupdateAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDentalStoreInsertorupdateAPIResponse
func GetAlibabaAlihealthDentalStoreInsertorupdateAPIResponse() *AlibabaAlihealthDentalStoreInsertorupdateAPIResponse {
	return poolAlibabaAlihealthDentalStoreInsertorupdateAPIResponse.Get().(*AlibabaAlihealthDentalStoreInsertorupdateAPIResponse)
}

// ReleaseAlibabaAlihealthDentalStoreInsertorupdateAPIResponse 将 AlibabaAlihealthDentalStoreInsertorupdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDentalStoreInsertorupdateAPIResponse(v *AlibabaAlihealthDentalStoreInsertorupdateAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDentalStoreInsertorupdateAPIResponse.Put(v)
}
