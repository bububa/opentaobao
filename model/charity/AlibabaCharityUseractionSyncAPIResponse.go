package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityUseractionSyncAPIResponse 用户公益行为同步 API返回值
// alibaba.charity.useraction.sync
//
// 外部公益活动，用户公益行为同步
type AlibabaCharityUseractionSyncAPIResponse struct {
	model.CommonResponse
	AlibabaCharityUseractionSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCharityUseractionSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCharityUseractionSyncAPIResponseModel).Reset()
}

// AlibabaCharityUseractionSyncAPIResponseModel is 用户公益行为同步 成功返回结果
type AlibabaCharityUseractionSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_useraction_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ThreehoursResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCharityUseractionSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCharityUseractionSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCharityUseractionSyncAPIResponse)
	},
}

// GetAlibabaCharityUseractionSyncAPIResponse 从 sync.Pool 获取 AlibabaCharityUseractionSyncAPIResponse
func GetAlibabaCharityUseractionSyncAPIResponse() *AlibabaCharityUseractionSyncAPIResponse {
	return poolAlibabaCharityUseractionSyncAPIResponse.Get().(*AlibabaCharityUseractionSyncAPIResponse)
}

// ReleaseAlibabaCharityUseractionSyncAPIResponse 将 AlibabaCharityUseractionSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCharityUseractionSyncAPIResponse(v *AlibabaCharityUseractionSyncAPIResponse) {
	v.Reset()
	poolAlibabaCharityUseractionSyncAPIResponse.Put(v)
}
