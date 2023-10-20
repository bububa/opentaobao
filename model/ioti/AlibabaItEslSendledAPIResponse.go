package ioti

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItEslSendledAPIResponse 厂测LED控制 API返回值
// alibaba.it.esl.sendled
//
// 针对厂测生产的的价签，增加led闪灯的接口，进行led 闪灯测试
type AlibabaItEslSendledAPIResponse struct {
	model.CommonResponse
	AlibabaItEslSendledAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaItEslSendledAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaItEslSendledAPIResponseModel).Reset()
}

// AlibabaItEslSendledAPIResponseModel is 厂测LED控制 成功返回结果
type AlibabaItEslSendledAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_it_esl_sendled_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Can not find Corresponding AP MAC with ESL
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaItEslSendledAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaItEslSendledAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaItEslSendledAPIResponse)
	},
}

// GetAlibabaItEslSendledAPIResponse 从 sync.Pool 获取 AlibabaItEslSendledAPIResponse
func GetAlibabaItEslSendledAPIResponse() *AlibabaItEslSendledAPIResponse {
	return poolAlibabaItEslSendledAPIResponse.Get().(*AlibabaItEslSendledAPIResponse)
}

// ReleaseAlibabaItEslSendledAPIResponse 将 AlibabaItEslSendledAPIResponse 保存到 sync.Pool
func ReleaseAlibabaItEslSendledAPIResponse(v *AlibabaItEslSendledAPIResponse) {
	v.Reset()
	poolAlibabaItEslSendledAPIResponse.Put(v)
}
