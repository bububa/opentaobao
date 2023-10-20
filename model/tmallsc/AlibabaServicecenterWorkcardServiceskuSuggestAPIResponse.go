package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse 服务商反馈需要履行的服务项 API返回值
// alibaba.servicecenter.workcard.servicesku.suggest
//
// 服务商反馈需要履行的服务项
type AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterWorkcardServiceskuSuggestAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterWorkcardServiceskuSuggestAPIResponseModel).Reset()
}

// AlibabaServicecenterWorkcardServiceskuSuggestAPIResponseModel is 服务商反馈需要履行的服务项 成功返回结果
type AlibabaServicecenterWorkcardServiceskuSuggestAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_workcard_servicesku_suggest_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardServiceskuSuggestAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.IsSuccess = false
}

var poolAlibabaServicecenterWorkcardServiceskuSuggestAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse)
	},
}

// GetAlibabaServicecenterWorkcardServiceskuSuggestAPIResponse 从 sync.Pool 获取 AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse
func GetAlibabaServicecenterWorkcardServiceskuSuggestAPIResponse() *AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse {
	return poolAlibabaServicecenterWorkcardServiceskuSuggestAPIResponse.Get().(*AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse)
}

// ReleaseAlibabaServicecenterWorkcardServiceskuSuggestAPIResponse 将 AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterWorkcardServiceskuSuggestAPIResponse(v *AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterWorkcardServiceskuSuggestAPIResponse.Put(v)
}
