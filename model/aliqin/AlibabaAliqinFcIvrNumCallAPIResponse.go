package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIvrNumCallAPIResponse ivr呼叫 API返回值
// alibaba.aliqin.fc.ivr.num.call
//
// ivr呼叫
type AlibabaAliqinFcIvrNumCallAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIvrNumCallAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIvrNumCallAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcIvrNumCallAPIResponseModel).Reset()
}

// AlibabaAliqinFcIvrNumCallAPIResponseModel is ivr呼叫 成功返回结果
type AlibabaAliqinFcIvrNumCallAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_ivr_num_call_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIvrNumCallResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIvrNumCallAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcIvrNumCallAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIvrNumCallAPIResponse)
	},
}

// GetAlibabaAliqinFcIvrNumCallAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcIvrNumCallAPIResponse
func GetAlibabaAliqinFcIvrNumCallAPIResponse() *AlibabaAliqinFcIvrNumCallAPIResponse {
	return poolAlibabaAliqinFcIvrNumCallAPIResponse.Get().(*AlibabaAliqinFcIvrNumCallAPIResponse)
}

// ReleaseAlibabaAliqinFcIvrNumCallAPIResponse 将 AlibabaAliqinFcIvrNumCallAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcIvrNumCallAPIResponse(v *AlibabaAliqinFcIvrNumCallAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcIvrNumCallAPIResponse.Put(v)
}
