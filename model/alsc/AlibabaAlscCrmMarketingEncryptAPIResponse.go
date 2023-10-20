package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmMarketingEncryptAPIResponse 加密 API返回值
// alibaba.alsc.crm.marketing.encrypt
//
// 加密
type AlibabaAlscCrmMarketingEncryptAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmMarketingEncryptAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmMarketingEncryptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmMarketingEncryptAPIResponseModel).Reset()
}

// AlibabaAlscCrmMarketingEncryptAPIResponseModel is 加密 成功返回结果
type AlibabaAlscCrmMarketingEncryptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_marketing_encrypt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmMarketingEncryptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmMarketingEncryptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmMarketingEncryptAPIResponse)
	},
}

// GetAlibabaAlscCrmMarketingEncryptAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmMarketingEncryptAPIResponse
func GetAlibabaAlscCrmMarketingEncryptAPIResponse() *AlibabaAlscCrmMarketingEncryptAPIResponse {
	return poolAlibabaAlscCrmMarketingEncryptAPIResponse.Get().(*AlibabaAlscCrmMarketingEncryptAPIResponse)
}

// ReleaseAlibabaAlscCrmMarketingEncryptAPIResponse 将 AlibabaAlscCrmMarketingEncryptAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmMarketingEncryptAPIResponse(v *AlibabaAlscCrmMarketingEncryptAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmMarketingEncryptAPIResponse.Put(v)
}
