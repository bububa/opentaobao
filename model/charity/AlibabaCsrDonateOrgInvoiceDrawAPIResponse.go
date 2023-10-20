package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrDonateOrgInvoiceDrawAPIResponse 机构开具商家票据信息 API返回值
// alibaba.csr.donate.org.invoice.draw
//
// 机构开具商家票据信息
type AlibabaCsrDonateOrgInvoiceDrawAPIResponse struct {
	model.CommonResponse
	AlibabaCsrDonateOrgInvoiceDrawAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCsrDonateOrgInvoiceDrawAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCsrDonateOrgInvoiceDrawAPIResponseModel).Reset()
}

// AlibabaCsrDonateOrgInvoiceDrawAPIResponseModel is 机构开具商家票据信息 成功返回结果
type AlibabaCsrDonateOrgInvoiceDrawAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_csr_donate_org_invoice_draw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应描述
	ResMsg string `json:"res_msg,omitempty" xml:"res_msg,omitempty"`
	// 空
	ResData string `json:"res_data,omitempty" xml:"res_data,omitempty"`
	// 无
	ResLocalizedMsg string `json:"res_localized_msg,omitempty" xml:"res_localized_msg,omitempty"`
	// 响应码
	ResCode int64 `json:"res_code,omitempty" xml:"res_code,omitempty"`
	// 成功
	ResSuccess bool `json:"res_success,omitempty" xml:"res_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCsrDonateOrgInvoiceDrawAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResMsg = ""
	m.ResData = ""
	m.ResLocalizedMsg = ""
	m.ResCode = 0
	m.ResSuccess = false
}

var poolAlibabaCsrDonateOrgInvoiceDrawAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCsrDonateOrgInvoiceDrawAPIResponse)
	},
}

// GetAlibabaCsrDonateOrgInvoiceDrawAPIResponse 从 sync.Pool 获取 AlibabaCsrDonateOrgInvoiceDrawAPIResponse
func GetAlibabaCsrDonateOrgInvoiceDrawAPIResponse() *AlibabaCsrDonateOrgInvoiceDrawAPIResponse {
	return poolAlibabaCsrDonateOrgInvoiceDrawAPIResponse.Get().(*AlibabaCsrDonateOrgInvoiceDrawAPIResponse)
}

// ReleaseAlibabaCsrDonateOrgInvoiceDrawAPIResponse 将 AlibabaCsrDonateOrgInvoiceDrawAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCsrDonateOrgInvoiceDrawAPIResponse(v *AlibabaCsrDonateOrgInvoiceDrawAPIResponse) {
	v.Reset()
	poolAlibabaCsrDonateOrgInvoiceDrawAPIResponse.Put(v)
}
