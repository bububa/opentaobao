package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscConsumerecordSyncAPIResponse 外域订单同步本地生活消费记录 API返回值
// alibaba.alsc.consumerecord.sync
//
// 外部第三方将本地生活app端下单数据同步到本地生活消费记录中心
type AlibabaAlscConsumerecordSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlscConsumerecordSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscConsumerecordSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscConsumerecordSyncAPIResponseModel).Reset()
}

// AlibabaAlscConsumerecordSyncAPIResponseModel is 外域订单同步本地生活消费记录 成功返回结果
type AlibabaAlscConsumerecordSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_consumerecord_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步返回结果
	Result *SingleDataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscConsumerecordSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscConsumerecordSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscConsumerecordSyncAPIResponse)
	},
}

// GetAlibabaAlscConsumerecordSyncAPIResponse 从 sync.Pool 获取 AlibabaAlscConsumerecordSyncAPIResponse
func GetAlibabaAlscConsumerecordSyncAPIResponse() *AlibabaAlscConsumerecordSyncAPIResponse {
	return poolAlibabaAlscConsumerecordSyncAPIResponse.Get().(*AlibabaAlscConsumerecordSyncAPIResponse)
}

// ReleaseAlibabaAlscConsumerecordSyncAPIResponse 将 AlibabaAlscConsumerecordSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscConsumerecordSyncAPIResponse(v *AlibabaAlscConsumerecordSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlscConsumerecordSyncAPIResponse.Put(v)
}
