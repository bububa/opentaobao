package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse 判断支付宝用户是否绑定淘宝账号 API返回值
// alibaba.aliqin.flow.alipay.isbindingtbaccount
//
// 判断支付宝用户是否绑定淘宝账号
type AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponseModel).Reset()
}

// AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponseModel is 判断支付宝用户是否绑定淘宝账号 成功返回结果
type AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_alipay_isbindingtbaccount_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// TRUE代表绑定，FALSE代表未绑定
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// errorCode
	AlicomFlowErrorCode string `json:"alicom_flow_error_code,omitempty" xml:"alicom_flow_error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// error
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = ""
	m.AlicomFlowErrorCode = ""
	m.ErrorMsg = ""
	m.Error = false
}

var poolAlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse)
	},
}

// GetAlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse 从 sync.Pool 获取 AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse
func GetAlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse() *AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse {
	return poolAlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse.Get().(*AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse)
}

// ReleaseAlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse 将 AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse(v *AlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFlowAlipayIsbindingtbaccountAPIResponse.Put(v)
}
