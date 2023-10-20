package fundplatform

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardorderMakeSuccessAPIResponse 通知制卡成功 API返回值
// alibaba.fundplatform.cardorder.make.success
//
// 当外部业务方调用资金平台异步制卡接口后，资金平台制卡成功后通知异步通知业务方
type AlibabaFundplatformCardorderMakeSuccessAPIResponse struct {
	model.CommonResponse
	AlibabaFundplatformCardorderMakeSuccessAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFundplatformCardorderMakeSuccessAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFundplatformCardorderMakeSuccessAPIResponseModel).Reset()
}

// AlibabaFundplatformCardorderMakeSuccessAPIResponseModel is 通知制卡成功 成功返回结果
type AlibabaFundplatformCardorderMakeSuccessAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_cardorder_make_success_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误详情
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 错误CODE
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFundplatformCardorderMakeSuccessAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMessage = ""
	m.ResultCode = ""
	m.Success = false
}

var poolAlibabaFundplatformCardorderMakeSuccessAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFundplatformCardorderMakeSuccessAPIResponse)
	},
}

// GetAlibabaFundplatformCardorderMakeSuccessAPIResponse 从 sync.Pool 获取 AlibabaFundplatformCardorderMakeSuccessAPIResponse
func GetAlibabaFundplatformCardorderMakeSuccessAPIResponse() *AlibabaFundplatformCardorderMakeSuccessAPIResponse {
	return poolAlibabaFundplatformCardorderMakeSuccessAPIResponse.Get().(*AlibabaFundplatformCardorderMakeSuccessAPIResponse)
}

// ReleaseAlibabaFundplatformCardorderMakeSuccessAPIResponse 将 AlibabaFundplatformCardorderMakeSuccessAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFundplatformCardorderMakeSuccessAPIResponse(v *AlibabaFundplatformCardorderMakeSuccessAPIResponse) {
	v.Reset()
	poolAlibabaFundplatformCardorderMakeSuccessAPIResponse.Put(v)
}
