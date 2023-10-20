package seaking

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingAuthmachineapiAPIResponse 机翻Api授权 API返回值
// alibaba.seaking.authmachineapi
//
// 机翻Api授权
type AlibabaSeakingAuthmachineapiAPIResponse struct {
	model.CommonResponse
	AlibabaSeakingAuthmachineapiAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSeakingAuthmachineapiAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSeakingAuthmachineapiAPIResponseModel).Reset()
}

// AlibabaSeakingAuthmachineapiAPIResponseModel is 机翻Api授权 成功返回结果
type AlibabaSeakingAuthmachineapiAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_authmachineapi_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSeakingAuthmachineapiAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaSeakingAuthmachineapiAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSeakingAuthmachineapiAPIResponse)
	},
}

// GetAlibabaSeakingAuthmachineapiAPIResponse 从 sync.Pool 获取 AlibabaSeakingAuthmachineapiAPIResponse
func GetAlibabaSeakingAuthmachineapiAPIResponse() *AlibabaSeakingAuthmachineapiAPIResponse {
	return poolAlibabaSeakingAuthmachineapiAPIResponse.Get().(*AlibabaSeakingAuthmachineapiAPIResponse)
}

// ReleaseAlibabaSeakingAuthmachineapiAPIResponse 将 AlibabaSeakingAuthmachineapiAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSeakingAuthmachineapiAPIResponse(v *AlibabaSeakingAuthmachineapiAPIResponse) {
	v.Reset()
	poolAlibabaSeakingAuthmachineapiAPIResponse.Put(v)
}
