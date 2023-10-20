package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacsrdonateorginvoiceundrawlistAPIResponse 获取机构待开票列表 API返回值
// alibaba.csr.donate.org.invoice.undraw.list
//
// 获取机构待开票列表
type AlibabacsrdonateorginvoiceundrawlistAPIResponse struct {
	model.CommonResponse
	AlibabacsrdonateorginvoiceundrawlistAPIResponseModel
}

// AlibabacsrdonateorginvoiceundrawlistAPIResponseModel is 获取机构待开票列表 成功返回结果
type AlibabacsrdonateorginvoiceundrawlistAPIResponseModel struct {
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
