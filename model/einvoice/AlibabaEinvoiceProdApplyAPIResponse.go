package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceProdApplyAPIResponse 提交发票申请 API返回值
// alibaba.einvoice.prod.apply
//
// 提交开票申请，如果商户授权自动开票则自动转开票，否则等待商户审核。
type AlibabaEinvoiceProdApplyAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceProdApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceProdApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceProdApplyAPIResponseModel).Reset()
}

// AlibabaEinvoiceProdApplyAPIResponseModel is 提交发票申请 成功返回结果
type AlibabaEinvoiceProdApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_prod_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceProdApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceProdApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceProdApplyAPIResponse)
	},
}

// GetAlibabaEinvoiceProdApplyAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceProdApplyAPIResponse
func GetAlibabaEinvoiceProdApplyAPIResponse() *AlibabaEinvoiceProdApplyAPIResponse {
	return poolAlibabaEinvoiceProdApplyAPIResponse.Get().(*AlibabaEinvoiceProdApplyAPIResponse)
}

// ReleaseAlibabaEinvoiceProdApplyAPIResponse 将 AlibabaEinvoiceProdApplyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceProdApplyAPIResponse(v *AlibabaEinvoiceProdApplyAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceProdApplyAPIResponse.Put(v)
}
