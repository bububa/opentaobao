package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse 根据健康ID获取支付宝ID API返回值
// alibaba.alihealth.docbase.userinfo.alipayid.get
//
// 根据健康ID获取支付宝ID
type AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponseModel).Reset()
}

// AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponseModel is 根据健康ID获取支付宝ID 成功返回结果
type AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_docbase_userinfo_alipayid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse)
	},
}

// GetAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse
func GetAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse() *AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse {
	return poolAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse.Get().(*AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse)
}

// ReleaseAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse 将 AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse(v *AlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDocbaseUserinfoAlipayidGetAPIResponse.Put(v)
}
