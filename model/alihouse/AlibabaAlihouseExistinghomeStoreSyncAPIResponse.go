package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreSyncAPIResponse 二手房标准门店数据同步 API返回值
// alibaba.alihouse.existinghome.store.sync
//
// 二手房标准门店数据同步
type AlibabaAlihouseExistinghomeStoreSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeStoreSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeStoreSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeStoreSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeStoreSyncAPIResponseModel is 二手房标准门店数据同步 成功返回结果
type AlibabaAlihouseExistinghomeStoreSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_store_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeStoreSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeStoreSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeStoreSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeStoreSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeStoreSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeStoreSyncAPIResponse
func GetAlibabaAlihouseExistinghomeStoreSyncAPIResponse() *AlibabaAlihouseExistinghomeStoreSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeStoreSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeStoreSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeStoreSyncAPIResponse 将 AlibabaAlihouseExistinghomeStoreSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeStoreSyncAPIResponse(v *AlibabaAlihouseExistinghomeStoreSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeStoreSyncAPIResponse.Put(v)
}
