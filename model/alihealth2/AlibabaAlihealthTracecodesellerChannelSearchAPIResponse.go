package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerChannelSearchAPIResponse 查询渠道商api API返回值
// alibaba.alihealth.tracecodeseller.channel.search
//
// 查询渠道商api
type AlibabaAlihealthTracecodesellerChannelSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesellerChannelSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerChannelSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodesellerChannelSearchAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodesellerChannelSearchAPIResponseModel is 查询渠道商api 成功返回结果
type AlibabaAlihealthTracecodesellerChannelSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_channel_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerChannelSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthTracecodesellerChannelSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerChannelSearchAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodesellerChannelSearchAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerChannelSearchAPIResponse
func GetAlibabaAlihealthTracecodesellerChannelSearchAPIResponse() *AlibabaAlihealthTracecodesellerChannelSearchAPIResponse {
	return poolAlibabaAlihealthTracecodesellerChannelSearchAPIResponse.Get().(*AlibabaAlihealthTracecodesellerChannelSearchAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodesellerChannelSearchAPIResponse 将 AlibabaAlihealthTracecodesellerChannelSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerChannelSearchAPIResponse(v *AlibabaAlihealthTracecodesellerChannelSearchAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerChannelSearchAPIResponse.Put(v)
}
