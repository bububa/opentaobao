package alink

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkMessageConfigListAPIResponse 查询消息开关配置列表 API返回值
// alibaba.alink.message.config.list
//
// 阿里智能获取消息开关配置列表
type AlibabaAlinkMessageConfigListAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkMessageConfigListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlinkMessageConfigListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlinkMessageConfigListAPIResponseModel).Reset()
}

// AlibabaAlinkMessageConfigListAPIResponseModel is 查询消息开关配置列表 成功返回结果
type AlibabaAlinkMessageConfigListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_message_config_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlinkMessageConfigListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlinkMessageConfigListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlinkMessageConfigListAPIResponse)
	},
}

// GetAlibabaAlinkMessageConfigListAPIResponse 从 sync.Pool 获取 AlibabaAlinkMessageConfigListAPIResponse
func GetAlibabaAlinkMessageConfigListAPIResponse() *AlibabaAlinkMessageConfigListAPIResponse {
	return poolAlibabaAlinkMessageConfigListAPIResponse.Get().(*AlibabaAlinkMessageConfigListAPIResponse)
}

// ReleaseAlibabaAlinkMessageConfigListAPIResponse 将 AlibabaAlinkMessageConfigListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlinkMessageConfigListAPIResponse(v *AlibabaAlinkMessageConfigListAPIResponse) {
	v.Reset()
	poolAlibabaAlinkMessageConfigListAPIResponse.Put(v)
}
