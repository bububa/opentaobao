package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceSettingRuleAPIResponse 发票规则设置 API返回值
// alitrip.btrip.invoice.setting.rule
//
// 发票规则设置
type AlitripBtripInvoiceSettingRuleAPIResponse struct {
	model.CommonResponse
	AlitripBtripInvoiceSettingRuleAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripInvoiceSettingRuleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripInvoiceSettingRuleAPIResponseModel).Reset()
}

// AlitripBtripInvoiceSettingRuleAPIResponseModel is 发票规则设置 成功返回结果
type AlitripBtripInvoiceSettingRuleAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_invoice_setting_rule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回值
	Result *OpenInvoiceRuleRs `json:"result,omitempty" xml:"result,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripInvoiceSettingRuleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Result = nil
	m.ResultCode = 0
}

var poolAlitripBtripInvoiceSettingRuleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripInvoiceSettingRuleAPIResponse)
	},
}

// GetAlitripBtripInvoiceSettingRuleAPIResponse 从 sync.Pool 获取 AlitripBtripInvoiceSettingRuleAPIResponse
func GetAlitripBtripInvoiceSettingRuleAPIResponse() *AlitripBtripInvoiceSettingRuleAPIResponse {
	return poolAlitripBtripInvoiceSettingRuleAPIResponse.Get().(*AlitripBtripInvoiceSettingRuleAPIResponse)
}

// ReleaseAlitripBtripInvoiceSettingRuleAPIResponse 将 AlitripBtripInvoiceSettingRuleAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripInvoiceSettingRuleAPIResponse(v *AlitripBtripInvoiceSettingRuleAPIResponse) {
	v.Reset()
	poolAlitripBtripInvoiceSettingRuleAPIResponse.Put(v)
}
