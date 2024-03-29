package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseChangeStoreTypeAPIResponse 融合店迁移门店 API返回值
// alibaba.alihouse.change.store.type
//
// 融合店迁移门店
type AlibabaAlihouseChangeStoreTypeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseChangeStoreTypeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseChangeStoreTypeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseChangeStoreTypeAPIResponseModel).Reset()
}

// AlibabaAlihouseChangeStoreTypeAPIResponseModel is 融合店迁移门店 成功返回结果
type AlibabaAlihouseChangeStoreTypeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_change_store_type_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 1
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 1
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 1
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseChangeStoreTypeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnCode = ""
	m.ReturnMessage = ""
	m.Data = false
	m.ReturnSuccess = false
}

var poolAlibabaAlihouseChangeStoreTypeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseChangeStoreTypeAPIResponse)
	},
}

// GetAlibabaAlihouseChangeStoreTypeAPIResponse 从 sync.Pool 获取 AlibabaAlihouseChangeStoreTypeAPIResponse
func GetAlibabaAlihouseChangeStoreTypeAPIResponse() *AlibabaAlihouseChangeStoreTypeAPIResponse {
	return poolAlibabaAlihouseChangeStoreTypeAPIResponse.Get().(*AlibabaAlihouseChangeStoreTypeAPIResponse)
}

// ReleaseAlibabaAlihouseChangeStoreTypeAPIResponse 将 AlibabaAlihouseChangeStoreTypeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseChangeStoreTypeAPIResponse(v *AlibabaAlihouseChangeStoreTypeAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseChangeStoreTypeAPIResponse.Put(v)
}
