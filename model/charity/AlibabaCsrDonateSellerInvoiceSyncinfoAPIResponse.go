package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse 链上同步商家票据信息 API返回值
// alibaba.csr.donate.seller.invoice.syncinfo
//
// 链上同步商家票据信息
type AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse struct {
	model.CommonResponse
	AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponseModel).Reset()
}

// AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponseModel is 链上同步商家票据信息 成功返回结果
type AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_csr_donate_seller_invoice_syncinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应描述信息
	ResMsg string `json:"res_msg,omitempty" xml:"res_msg,omitempty"`
	// 响应数据
	ResData string `json:"res_data,omitempty" xml:"res_data,omitempty"`
	// 忽略
	ResLocalizedMsg string `json:"res_localized_msg,omitempty" xml:"res_localized_msg,omitempty"`
	// 响应码
	ResCode int64 `json:"res_code,omitempty" xml:"res_code,omitempty"`
	// 是否处理成功
	ResSuccess bool `json:"res_success,omitempty" xml:"res_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResMsg = ""
	m.ResData = ""
	m.ResLocalizedMsg = ""
	m.ResCode = 0
	m.ResSuccess = false
}

var poolAlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse)
	},
}

// GetAlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse 从 sync.Pool 获取 AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse
func GetAlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse() *AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse {
	return poolAlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse.Get().(*AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse)
}

// ReleaseAlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse 将 AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse(v *AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse) {
	v.Reset()
	poolAlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse.Put(v)
}
