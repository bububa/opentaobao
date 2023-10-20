package alink

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkMessageConfigSetAPIResponse 消息提醒开关 API返回值
// alibaba.alink.message.config.set
//
// 阿里智能消息开关
type AlibabaAlinkMessageConfigSetAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkMessageConfigSetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlinkMessageConfigSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlinkMessageConfigSetAPIResponseModel).Reset()
}

// AlibabaAlinkMessageConfigSetAPIResponseModel is 消息提醒开关 成功返回结果
type AlibabaAlinkMessageConfigSetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_message_config_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlinkMessageConfigSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlinkMessageConfigSetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlinkMessageConfigSetAPIResponse)
	},
}

// GetAlibabaAlinkMessageConfigSetAPIResponse 从 sync.Pool 获取 AlibabaAlinkMessageConfigSetAPIResponse
func GetAlibabaAlinkMessageConfigSetAPIResponse() *AlibabaAlinkMessageConfigSetAPIResponse {
	return poolAlibabaAlinkMessageConfigSetAPIResponse.Get().(*AlibabaAlinkMessageConfigSetAPIResponse)
}

// ReleaseAlibabaAlinkMessageConfigSetAPIResponse 将 AlibabaAlinkMessageConfigSetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlinkMessageConfigSetAPIResponse(v *AlibabaAlinkMessageConfigSetAPIResponse) {
	v.Reset()
	poolAlibabaAlinkMessageConfigSetAPIResponse.Put(v)
}
