package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderPrePayActionAPIResponse 服务商预付款完成接口 API返回值
// alibaba.idle.tender.pre.pay.action
//
// 服务商预付款完成接口
type AlibabaIdleTenderPrePayActionAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTenderPrePayActionAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleTenderPrePayActionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleTenderPrePayActionAPIResponseModel).Reset()
}

// AlibabaIdleTenderPrePayActionAPIResponseModel is 服务商预付款完成接口 成功返回结果
type AlibabaIdleTenderPrePayActionAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_pre_pay_action_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleTenderPrePayActionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleTenderPrePayActionAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTenderPrePayActionAPIResponse)
	},
}

// GetAlibabaIdleTenderPrePayActionAPIResponse 从 sync.Pool 获取 AlibabaIdleTenderPrePayActionAPIResponse
func GetAlibabaIdleTenderPrePayActionAPIResponse() *AlibabaIdleTenderPrePayActionAPIResponse {
	return poolAlibabaIdleTenderPrePayActionAPIResponse.Get().(*AlibabaIdleTenderPrePayActionAPIResponse)
}

// ReleaseAlibabaIdleTenderPrePayActionAPIResponse 将 AlibabaIdleTenderPrePayActionAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleTenderPrePayActionAPIResponse(v *AlibabaIdleTenderPrePayActionAPIResponse) {
	v.Reset()
	poolAlibabaIdleTenderPrePayActionAPIResponse.Put(v)
}
