package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmVoucherStatusChangeAPIResponse 优惠券状态更改 API返回值
// alibaba.alsc.crm.voucher.status.change
//
// 核销优惠券
type AlibabaAlscCrmVoucherStatusChangeAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmVoucherStatusChangeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmVoucherStatusChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmVoucherStatusChangeAPIResponseModel).Reset()
}

// AlibabaAlscCrmVoucherStatusChangeAPIResponseModel is 优惠券状态更改 成功返回结果
type AlibabaAlscCrmVoucherStatusChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_voucher_status_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmVoucherStatusChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmVoucherStatusChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmVoucherStatusChangeAPIResponse)
	},
}

// GetAlibabaAlscCrmVoucherStatusChangeAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmVoucherStatusChangeAPIResponse
func GetAlibabaAlscCrmVoucherStatusChangeAPIResponse() *AlibabaAlscCrmVoucherStatusChangeAPIResponse {
	return poolAlibabaAlscCrmVoucherStatusChangeAPIResponse.Get().(*AlibabaAlscCrmVoucherStatusChangeAPIResponse)
}

// ReleaseAlibabaAlscCrmVoucherStatusChangeAPIResponse 将 AlibabaAlscCrmVoucherStatusChangeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmVoucherStatusChangeAPIResponse(v *AlibabaAlscCrmVoucherStatusChangeAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmVoucherStatusChangeAPIResponse.Put(v)
}
