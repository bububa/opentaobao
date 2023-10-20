package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomVtDistributeSendcodeAPIResponse 通信业务外放发送验证码 API返回值
// alibaba.alicom.vt.distribute.sendcode
//
// 通信业务外放发送验证码
type AlibabaAlicomVtDistributeSendcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlicomVtDistributeSendcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlicomVtDistributeSendcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlicomVtDistributeSendcodeAPIResponseModel).Reset()
}

// AlibabaAlicomVtDistributeSendcodeAPIResponseModel is 通信业务外放发送验证码 成功返回结果
type AlibabaAlicomVtDistributeSendcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_vt_distribute_sendcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlicomVtDistributeSendcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlicomVtDistributeSendcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlicomVtDistributeSendcodeAPIResponse)
	},
}

// GetAlibabaAlicomVtDistributeSendcodeAPIResponse 从 sync.Pool 获取 AlibabaAlicomVtDistributeSendcodeAPIResponse
func GetAlibabaAlicomVtDistributeSendcodeAPIResponse() *AlibabaAlicomVtDistributeSendcodeAPIResponse {
	return poolAlibabaAlicomVtDistributeSendcodeAPIResponse.Get().(*AlibabaAlicomVtDistributeSendcodeAPIResponse)
}

// ReleaseAlibabaAlicomVtDistributeSendcodeAPIResponse 将 AlibabaAlicomVtDistributeSendcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlicomVtDistributeSendcodeAPIResponse(v *AlibabaAlicomVtDistributeSendcodeAPIResponse) {
	v.Reset()
	poolAlibabaAlicomVtDistributeSendcodeAPIResponse.Put(v)
}
