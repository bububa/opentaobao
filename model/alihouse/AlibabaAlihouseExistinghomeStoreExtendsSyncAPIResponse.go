package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse 门店扩展信息变更 API返回值
// alibaba.alihouse.existinghome.store.extends.sync
//
// 门店扩展信息变更
type AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponseModel is 门店扩展信息变更 成功返回结果
type AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_store_extends_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// aaa
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// aaa
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// aaa
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// aaa
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnCode = ""
	m.ReturnMessage = ""
	m.Data = false
	m.ReturnSuccess = false
}

var poolAlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse
func GetAlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse() *AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse 将 AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse(v *AlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeStoreExtendsSyncAPIResponse.Put(v)
}
