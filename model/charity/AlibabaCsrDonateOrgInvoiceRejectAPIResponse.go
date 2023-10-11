package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrDonateOrgInvoiceRejectAPIResponse 机构驳回商家票据信息 API返回值
// alibaba.csr.donate.org.invoice.reject
//
// 机构驳回商家票据信息
type AlibabaCsrDonateOrgInvoiceRejectAPIResponse struct {
	model.CommonResponse
	AlibabaCsrDonateOrgInvoiceRejectAPIResponseModel
}

// AlibabaCsrDonateOrgInvoiceRejectAPIResponseModel is 机构驳回商家票据信息 成功返回结果
type AlibabaCsrDonateOrgInvoiceRejectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_csr_donate_org_invoice_reject_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应描述
	ResMsg string `json:"res_msg,omitempty" xml:"res_msg,omitempty"`
	// 无
	ResData string `json:"res_data,omitempty" xml:"res_data,omitempty"`
	// 无
	ResLocalizedMsg string `json:"res_localized_msg,omitempty" xml:"res_localized_msg,omitempty"`
	// 响应码
	ResCode int64 `json:"res_code,omitempty" xml:"res_code,omitempty"`
	// 成功
	ResSuccess bool `json:"res_success,omitempty" xml:"res_success,omitempty"`
}
