package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerVoucherListAPIResponse 获取顾客优惠券列表 API返回值
// alibaba.alsc.crm.customer.voucher.list
//
// 获取顾客优惠券列表
type AlibabaAlscCrmCustomerVoucherListAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCustomerVoucherListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCustomerVoucherListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmCustomerVoucherListAPIResponseModel).Reset()
}

// AlibabaAlscCrmCustomerVoucherListAPIResponseModel is 获取顾客优惠券列表 成功返回结果
type AlibabaAlscCrmCustomerVoucherListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_customer_voucher_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmCustomerVoucherListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmCustomerVoucherListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmCustomerVoucherListAPIResponse)
	},
}

// GetAlibabaAlscCrmCustomerVoucherListAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmCustomerVoucherListAPIResponse
func GetAlibabaAlscCrmCustomerVoucherListAPIResponse() *AlibabaAlscCrmCustomerVoucherListAPIResponse {
	return poolAlibabaAlscCrmCustomerVoucherListAPIResponse.Get().(*AlibabaAlscCrmCustomerVoucherListAPIResponse)
}

// ReleaseAlibabaAlscCrmCustomerVoucherListAPIResponse 将 AlibabaAlscCrmCustomerVoucherListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmCustomerVoucherListAPIResponse(v *AlibabaAlscCrmCustomerVoucherListAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmCustomerVoucherListAPIResponse.Put(v)
}
