package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmVoucherSendAPIResponse 发送券给指定用户 API返回值
// alibaba.alsc.crm.voucher.send
//
// 发送券给指定用户
type AlibabaAlscCrmVoucherSendAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmVoucherSendAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmVoucherSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmVoucherSendAPIResponseModel).Reset()
}

// AlibabaAlscCrmVoucherSendAPIResponseModel is 发送券给指定用户 成功返回结果
type AlibabaAlscCrmVoucherSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_voucher_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmVoucherSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmVoucherSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmVoucherSendAPIResponse)
	},
}

// GetAlibabaAlscCrmVoucherSendAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmVoucherSendAPIResponse
func GetAlibabaAlscCrmVoucherSendAPIResponse() *AlibabaAlscCrmVoucherSendAPIResponse {
	return poolAlibabaAlscCrmVoucherSendAPIResponse.Get().(*AlibabaAlscCrmVoucherSendAPIResponse)
}

// ReleaseAlibabaAlscCrmVoucherSendAPIResponse 将 AlibabaAlscCrmVoucherSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmVoucherSendAPIResponse(v *AlibabaAlscCrmVoucherSendAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmVoucherSendAPIResponse.Put(v)
}
