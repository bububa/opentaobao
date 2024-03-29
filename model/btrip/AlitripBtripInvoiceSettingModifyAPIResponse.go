package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceSettingModifyAPIResponse 发票变更 API返回值
// alitrip.btrip.invoice.setting.modify
//
// 发票变更
type AlitripBtripInvoiceSettingModifyAPIResponse struct {
	model.CommonResponse
	AlitripBtripInvoiceSettingModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripInvoiceSettingModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripInvoiceSettingModifyAPIResponseModel).Reset()
}

// AlitripBtripInvoiceSettingModifyAPIResponseModel is 发票变更 成功返回结果
type AlitripBtripInvoiceSettingModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_invoice_setting_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 请求是否成功
	SuccessFlag bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripInvoiceSettingModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = 0
	m.Result = false
	m.SuccessFlag = false
}

var poolAlitripBtripInvoiceSettingModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripInvoiceSettingModifyAPIResponse)
	},
}

// GetAlitripBtripInvoiceSettingModifyAPIResponse 从 sync.Pool 获取 AlitripBtripInvoiceSettingModifyAPIResponse
func GetAlitripBtripInvoiceSettingModifyAPIResponse() *AlitripBtripInvoiceSettingModifyAPIResponse {
	return poolAlitripBtripInvoiceSettingModifyAPIResponse.Get().(*AlitripBtripInvoiceSettingModifyAPIResponse)
}

// ReleaseAlitripBtripInvoiceSettingModifyAPIResponse 将 AlitripBtripInvoiceSettingModifyAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripInvoiceSettingModifyAPIResponse(v *AlitripBtripInvoiceSettingModifyAPIResponse) {
	v.Reset()
	poolAlitripBtripInvoiceSettingModifyAPIResponse.Put(v)
}
