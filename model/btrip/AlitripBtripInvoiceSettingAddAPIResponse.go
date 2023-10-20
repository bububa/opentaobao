package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceSettingAddAPIResponse 发票设置 API返回值
// alitrip.btrip.invoice.setting.add
//
// 发票设置
type AlitripBtripInvoiceSettingAddAPIResponse struct {
	model.CommonResponse
	AlitripBtripInvoiceSettingAddAPIResponseModel
}

// AlitripBtripInvoiceSettingAddAPIResponseModel is 发票设置 成功返回结果
type AlitripBtripInvoiceSettingAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_invoice_setting_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果值
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 请求是否成功
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}
