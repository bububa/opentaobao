package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrDonateOrgInvoiceUndrawListAPIResponse 获取机构待开票列表 API返回值
// alibaba.csr.donate.org.invoice.undraw.list
//
// 获取机构待开票列表
type AlibabaCsrDonateOrgInvoiceUndrawListAPIResponse struct {
	model.CommonResponse
	AlibabaCsrDonateOrgInvoiceUndrawListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCsrDonateOrgInvoiceUndrawListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCsrDonateOrgInvoiceUndrawListAPIResponseModel).Reset()
}

// AlibabaCsrDonateOrgInvoiceUndrawListAPIResponseModel is 获取机构待开票列表 成功返回结果
type AlibabaCsrDonateOrgInvoiceUndrawListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_csr_donate_org_invoice_undraw_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应描述信息
	ResMsg string `json:"res_msg,omitempty" xml:"res_msg,omitempty"`
	// 无
	ResLocalizedMsg string `json:"res_localized_msg,omitempty" xml:"res_localized_msg,omitempty"`
	// 响应码，200成功，非200失败
	ResCode int64 `json:"res_code,omitempty" xml:"res_code,omitempty"`
	// 响应数据
	ResData *CsrPage `json:"res_data,omitempty" xml:"res_data,omitempty"`
	// 是否成功响应
	ResSuccess bool `json:"res_success,omitempty" xml:"res_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCsrDonateOrgInvoiceUndrawListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResMsg = ""
	m.ResLocalizedMsg = ""
	m.ResCode = 0
	m.ResData = nil
	m.ResSuccess = false
}

var poolAlibabaCsrDonateOrgInvoiceUndrawListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCsrDonateOrgInvoiceUndrawListAPIResponse)
	},
}

// GetAlibabaCsrDonateOrgInvoiceUndrawListAPIResponse 从 sync.Pool 获取 AlibabaCsrDonateOrgInvoiceUndrawListAPIResponse
func GetAlibabaCsrDonateOrgInvoiceUndrawListAPIResponse() *AlibabaCsrDonateOrgInvoiceUndrawListAPIResponse {
	return poolAlibabaCsrDonateOrgInvoiceUndrawListAPIResponse.Get().(*AlibabaCsrDonateOrgInvoiceUndrawListAPIResponse)
}

// ReleaseAlibabaCsrDonateOrgInvoiceUndrawListAPIResponse 将 AlibabaCsrDonateOrgInvoiceUndrawListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCsrDonateOrgInvoiceUndrawListAPIResponse(v *AlibabaCsrDonateOrgInvoiceUndrawListAPIResponse) {
	v.Reset()
	poolAlibabaCsrDonateOrgInvoiceUndrawListAPIResponse.Put(v)
}
