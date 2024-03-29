package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderDtdResendAPIResponse 门店自送重发码 API返回值
// taobao.omniorder.dtd.resend
//
// 该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次
type TaobaoOmniorderDtdResendAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderDtdResendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderDtdResendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderDtdResendAPIResponseModel).Reset()
}

// TaobaoOmniorderDtdResendAPIResponseModel is 门店自送重发码 成功返回结果
type TaobaoOmniorderDtdResendAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_dtd_resend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码，为0表示成功，非0表示失败
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderDtdResendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.Message = ""
}

var poolTaobaoOmniorderDtdResendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderDtdResendAPIResponse)
	},
}

// GetTaobaoOmniorderDtdResendAPIResponse 从 sync.Pool 获取 TaobaoOmniorderDtdResendAPIResponse
func GetTaobaoOmniorderDtdResendAPIResponse() *TaobaoOmniorderDtdResendAPIResponse {
	return poolTaobaoOmniorderDtdResendAPIResponse.Get().(*TaobaoOmniorderDtdResendAPIResponse)
}

// ReleaseTaobaoOmniorderDtdResendAPIResponse 将 TaobaoOmniorderDtdResendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderDtdResendAPIResponse(v *TaobaoOmniorderDtdResendAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderDtdResendAPIResponse.Put(v)
}
