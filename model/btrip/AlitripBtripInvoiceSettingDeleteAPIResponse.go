package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripInvoiceSettingDeleteAPIResponse 发票删除 API返回值
// alitrip.btrip.invoice.setting.delete
//
// 发票删除
type AlitripBtripInvoiceSettingDeleteAPIResponse struct {
	model.CommonResponse
	AlitripBtripInvoiceSettingDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripInvoiceSettingDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripInvoiceSettingDeleteAPIResponseModel).Reset()
}

// AlitripBtripInvoiceSettingDeleteAPIResponseModel is 发票删除 成功返回结果
type AlitripBtripInvoiceSettingDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_invoice_setting_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripInvoiceSettingDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = 0
	m.Result = false
}

var poolAlitripBtripInvoiceSettingDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripInvoiceSettingDeleteAPIResponse)
	},
}

// GetAlitripBtripInvoiceSettingDeleteAPIResponse 从 sync.Pool 获取 AlitripBtripInvoiceSettingDeleteAPIResponse
func GetAlitripBtripInvoiceSettingDeleteAPIResponse() *AlitripBtripInvoiceSettingDeleteAPIResponse {
	return poolAlitripBtripInvoiceSettingDeleteAPIResponse.Get().(*AlitripBtripInvoiceSettingDeleteAPIResponse)
}

// ReleaseAlitripBtripInvoiceSettingDeleteAPIResponse 将 AlitripBtripInvoiceSettingDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripInvoiceSettingDeleteAPIResponse(v *AlitripBtripInvoiceSettingDeleteAPIResponse) {
	v.Reset()
	poolAlitripBtripInvoiceSettingDeleteAPIResponse.Put(v)
}
