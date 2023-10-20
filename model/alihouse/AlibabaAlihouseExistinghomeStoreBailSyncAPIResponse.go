package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse 门店保证金余额同步 API返回值
// alibaba.alihouse.existinghome.store.bail.sync
//
// 门店保证金余额同步
type AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeStoreBailSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeStoreBailSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeStoreBailSyncAPIResponseModel is 门店保证金余额同步 成功返回结果
type AlibabaAlihouseExistinghomeStoreBailSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_store_bail_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// code
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// message
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 是否处理成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// success
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeStoreBailSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnCode = ""
	m.ReturnMessage = ""
	m.Data = false
	m.ReturnSuccess = false
}

var poolAlibabaAlihouseExistinghomeStoreBailSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeStoreBailSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse
func GetAlibabaAlihouseExistinghomeStoreBailSyncAPIResponse() *AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeStoreBailSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeStoreBailSyncAPIResponse 将 AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeStoreBailSyncAPIResponse(v *AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeStoreBailSyncAPIResponse.Put(v)
}
