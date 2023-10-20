package westcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmUserConsumerGetAPIResponse 获取指定用户的消费总额 API返回值
// alibaba.westcrm.user.consumer.get
//
// 获取指定用户的消费总额
type AlibabaWestcrmUserConsumerGetAPIResponse struct {
	model.CommonResponse
	AlibabaWestcrmUserConsumerGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWestcrmUserConsumerGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWestcrmUserConsumerGetAPIResponseModel).Reset()
}

// AlibabaWestcrmUserConsumerGetAPIResponseModel is 获取指定用户的消费总额 成功返回结果
type AlibabaWestcrmUserConsumerGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_westcrm_user_consumer_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWestcrmUserConsumerGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWestcrmUserConsumerGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWestcrmUserConsumerGetAPIResponse)
	},
}

// GetAlibabaWestcrmUserConsumerGetAPIResponse 从 sync.Pool 获取 AlibabaWestcrmUserConsumerGetAPIResponse
func GetAlibabaWestcrmUserConsumerGetAPIResponse() *AlibabaWestcrmUserConsumerGetAPIResponse {
	return poolAlibabaWestcrmUserConsumerGetAPIResponse.Get().(*AlibabaWestcrmUserConsumerGetAPIResponse)
}

// ReleaseAlibabaWestcrmUserConsumerGetAPIResponse 将 AlibabaWestcrmUserConsumerGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWestcrmUserConsumerGetAPIResponse(v *AlibabaWestcrmUserConsumerGetAPIResponse) {
	v.Reset()
	poolAlibabaWestcrmUserConsumerGetAPIResponse.Put(v)
}
