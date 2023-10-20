package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeSignatureSyncAPIResponse 二手房电子签章数据同步 API返回值
// alibaba.alihouse.existinghome.signature.sync
//
// 二手房电子签章数据同步
type AlibabaAlihouseExistinghomeSignatureSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeSignatureSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeSignatureSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeSignatureSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeSignatureSyncAPIResponseModel is 二手房电子签章数据同步 成功返回结果
type AlibabaAlihouseExistinghomeSignatureSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_signature_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeSignatureSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeSignatureSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeSignatureSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeSignatureSyncAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeSignatureSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeSignatureSyncAPIResponse
func GetAlibabaAlihouseExistinghomeSignatureSyncAPIResponse() *AlibabaAlihouseExistinghomeSignatureSyncAPIResponse {
	return poolAlibabaAlihouseExistinghomeSignatureSyncAPIResponse.Get().(*AlibabaAlihouseExistinghomeSignatureSyncAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeSignatureSyncAPIResponse 将 AlibabaAlihouseExistinghomeSignatureSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeSignatureSyncAPIResponse(v *AlibabaAlihouseExistinghomeSignatureSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeSignatureSyncAPIResponse.Put(v)
}
