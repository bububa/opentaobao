package zqs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaZqsFulfillCompleteAPIResponse 周期购履约完成接口 API返回值
// alibaba.zqs.fulfill.complete
//
// 周期购履约完成接口
type AlibabaZqsFulfillCompleteAPIResponse struct {
	model.CommonResponse
	AlibabaZqsFulfillCompleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaZqsFulfillCompleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaZqsFulfillCompleteAPIResponseModel).Reset()
}

// AlibabaZqsFulfillCompleteAPIResponseModel is 周期购履约完成接口 成功返回结果
type AlibabaZqsFulfillCompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_zqs_fulfill_complete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异常描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 异常code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 执行结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaZqsFulfillCompleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = ""
	m.Result = false
}

var poolAlibabaZqsFulfillCompleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaZqsFulfillCompleteAPIResponse)
	},
}

// GetAlibabaZqsFulfillCompleteAPIResponse 从 sync.Pool 获取 AlibabaZqsFulfillCompleteAPIResponse
func GetAlibabaZqsFulfillCompleteAPIResponse() *AlibabaZqsFulfillCompleteAPIResponse {
	return poolAlibabaZqsFulfillCompleteAPIResponse.Get().(*AlibabaZqsFulfillCompleteAPIResponse)
}

// ReleaseAlibabaZqsFulfillCompleteAPIResponse 将 AlibabaZqsFulfillCompleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaZqsFulfillCompleteAPIResponse(v *AlibabaZqsFulfillCompleteAPIResponse) {
	v.Reset()
	poolAlibabaZqsFulfillCompleteAPIResponse.Put(v)
}
